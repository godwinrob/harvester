// Galaxy types matching the backend API

export interface Galaxy {
  id: string
  name: string
  ownerUserID: string
  enabled: boolean
  dateCreated: string
  dateUpdated: string
}

export interface CreateGalaxyInput {
  name: string
  ownerUserID: string
}

export interface UpdateGalaxyInput {
  name?: string
  ownerUserID?: string
  enabled?: boolean
}

// Bulk operation types
export interface BulkCreateGalaxyInput {
  items: CreateGalaxyInput[]
}

export interface BulkUpdateGalaxyInput {
  items: Array<{
    id: string
    data: UpdateGalaxyInput
  }>
}

export interface BulkDeleteGalaxyInput {
  ids: string[]
}

// Query params
export interface GalaxyListParams {
  page?: number
  rows?: number
  orderBy?: {
    field: string
    direction: 'ASC' | 'DESC'
  }
  name?: string
}
