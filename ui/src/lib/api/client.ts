import { ApiError, NetworkError } from './errors'
import type { PaginationParams } from './types'
import { config } from '../config'

const API_BASE_URL = config.apiUrl

// Auth token storage - will be managed by auth store
let authToken: string | null = null

export function setAuthToken(token: string | null) {
  authToken = token
}

export function getAuthToken(): string | null {
  return authToken
}

interface RequestConfig extends Omit<RequestInit, 'body'> {
  body?: unknown
  params?: Record<string, string | number | boolean | undefined>
}

async function request<T>(
  endpoint: string,
  { body, params, ...customConfig }: RequestConfig = {}
): Promise<T> {
  const url = new URL(`${API_BASE_URL}${endpoint}`, window.location.origin)

  // Add query params
  if (params) {
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        url.searchParams.set(key, String(value))
      }
    })
  }

  const config: RequestInit = {
    method: body ? 'POST' : 'GET',
    ...customConfig,
    headers: {
      'Content-Type': 'application/json',
      ...customConfig.headers,
    },
  }

  if (body) {
    config.body = JSON.stringify(body)
  }

  // Add auth token when available
  if (authToken) {
    config.headers = {
      ...config.headers,
      Authorization: `Bearer ${authToken}`,
    }
  }

  try {
    const response = await fetch(url.toString(), config)

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}))
      throw ApiError.fromResponse(response.status, errorData)
    }

    // Handle 204 No Content
    if (response.status === 204) {
      return undefined as T
    }

    return response.json()
  } catch (error) {
    if (error instanceof ApiError) {
      throw error
    }
    throw new NetworkError('Network request failed', { cause: error })
  }
}

/**
 * Build pagination query parameters
 */
export function buildPaginationParams(
  params: PaginationParams
): Record<string, string | number> {
  const result: Record<string, string | number> = {}

  if (params.page !== undefined) {
    result.page = params.page
  }

  if (params.rows !== undefined) {
    result.rows = params.rows
  }

  if (params.orderBy) {
    result.orderBy = `${params.orderBy.field},${params.orderBy.direction}`
  }

  return result
}

/**
 * API client with typed methods
 */
export const api = {
  get: <T>(endpoint: string, params?: RequestConfig['params']) =>
    request<T>(endpoint, { params }),

  post: <T>(endpoint: string, body: unknown) =>
    request<T>(endpoint, { method: 'POST', body }),

  put: <T>(endpoint: string, body: unknown) =>
    request<T>(endpoint, { method: 'PUT', body }),

  patch: <T>(endpoint: string, body: unknown) =>
    request<T>(endpoint, { method: 'PATCH', body }),

  delete: <T>(endpoint: string, body?: unknown) =>
    request<T>(endpoint, { method: 'DELETE', body }),
}
