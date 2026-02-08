import type { ApiErrorResponse, BulkErrorItem } from './types'

export class ApiError extends Error {
  status: number
  code?: string
  fields?: Record<string, string[]>
  errors?: BulkErrorItem[]

  constructor(
    message: string,
    status: number,
    code?: string,
    fields?: Record<string, string[]>,
    errors?: BulkErrorItem[]
  ) {
    super(message)
    this.name = 'ApiError'
    this.status = status
    this.code = code
    this.fields = fields
    this.errors = errors
  }

  static fromResponse(status: number, data: unknown): ApiError {
    if (typeof data === 'object' && data !== null) {
      const errorData = data as ApiErrorResponse
      return new ApiError(
        errorData.error ?? 'An error occurred',
        status,
        errorData.code,
        errorData.fields,
        errorData.errors
      )
    }
    return new ApiError('An error occurred', status)
  }

  get isValidationError(): boolean {
    return this.status === 400 && (!!this.fields || !!this.errors)
  }

  get isUnauthorized(): boolean {
    return this.status === 401
  }

  get isForbidden(): boolean {
    return this.status === 403
  }

  get isNotFound(): boolean {
    return this.status === 404
  }

  get isConflict(): boolean {
    return this.status === 409
  }

  /**
   * Get field errors for form integration
   * Returns the first error message for each field
   */
  getFieldErrors(): Record<string, string> {
    if (!this.fields) return {}
    return Object.fromEntries(
      Object.entries(this.fields).map(([field, messages]) => [field, messages[0]])
    )
  }

  /**
   * Get bulk operation errors by index
   * Useful for displaying errors in bulk import UI
   */
  getBulkErrors(): Map<number, { message: string; fields?: Record<string, string> }> {
    const map = new Map<number, { message: string; fields?: Record<string, string> }>()
    if (this.errors) {
      this.errors.forEach(({ index, error, fields }) => {
        map.set(index, {
          message: error,
          fields: fields
            ? Object.fromEntries(
                Object.entries(fields).map(([f, m]) => [f, m[0]])
              )
            : undefined,
        })
      })
    }
    return map
  }
}

export class NetworkError extends Error {
  constructor(message: string, options?: ErrorOptions) {
    super(message, options)
    this.name = 'NetworkError'
  }
}
