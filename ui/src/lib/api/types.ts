// Common API types

export interface PaginationParams {
  page?: number
  rows?: number
  orderBy?: {
    field: string
    direction: 'ASC' | 'DESC'
  }
}

export interface PaginatedResponse<T> {
  items: T[]
  total: number
  page: number
  rowsPerPage: number
}

export interface BulkCreateResponse<T> {
  items: T[]
  created: number
}

export interface BulkUpdateResponse<T> {
  items: T[]
  updated: number
}

export interface BulkDeleteResponse {
  deleted: number
}

export interface BulkErrorItem {
  index: number
  field?: string
  error: string
}

export interface ApiErrorResponse {
  error?: string
  code?: string
  fields?: Record<string, string[]>
  errors?: BulkErrorItem[]
}
