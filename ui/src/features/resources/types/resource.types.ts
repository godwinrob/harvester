// Resource types matching the backend API

export interface Resource {
  id: string
  name: string
  galaxyID: string
  addedAtDate: string
  updatedAtDate: string
  addedUserID: string
  resourceType: string
  unavailableAt: string | null
  unavailableUserID: string | null
  verified: boolean
  verifiedUserID: string | null
  // Stats
  cr: number  // Cold Resistance
  cd: number  // Conductivity
  dr: number  // Decay Resistance
  fl: number  // Flavor
  hr: number  // Heat Resistance
  ma: number  // Malleability
  pe: number  // Potential Energy
  oq: number  // Overall Quality
  sr: number  // Shock Resistance
  ut: number  // Unit Toughness
  er: number  // Entangle Resistance
}

export interface CreateResourceInput {
  name: string
  galaxyID: string
  addedUserID: string
  resourceType: string
  // Stats (all optional, default to 0)
  cr?: number
  cd?: number
  dr?: number
  fl?: number
  hr?: number
  ma?: number
  pe?: number
  oq?: number
  sr?: number
  ut?: number
  er?: number
}

export interface UpdateResourceInput {
  name?: string
  galaxyID?: string
  resourceType?: string
  unavailableAt?: string | null
  unavailableUserID?: string | null
  verified?: boolean
  verifiedUserID?: string | null
  // Stats
  cr?: number
  cd?: number
  dr?: number
  fl?: number
  hr?: number
  ma?: number
  pe?: number
  oq?: number
  sr?: number
  ut?: number
  er?: number
}

// Bulk operation types
export interface BulkCreateResourceInput {
  items: CreateResourceInput[]
}

export interface BulkUpdateResourceInput {
  items: Array<{
    id: string
    data: UpdateResourceInput
  }>
}

export interface BulkDeleteResourceInput {
  ids: string[]
}

// Query params
export interface ResourceListParams {
  page?: number
  rows?: number
  orderBy?: {
    field: string
    direction: 'ASC' | 'DESC'
  }
  name?: string
  resourceType?: string
  resourceGroup?: string
  galaxyID?: string
}

// Stats metadata for UI
export const RESOURCE_STATS = [
  { key: 'cr', label: 'Cold Resistance', abbr: 'CR' },
  { key: 'cd', label: 'Conductivity', abbr: 'CD' },
  { key: 'dr', label: 'Decay Resistance', abbr: 'DR' },
  { key: 'fl', label: 'Flavor', abbr: 'FL' },
  { key: 'hr', label: 'Heat Resistance', abbr: 'HR' },
  { key: 'ma', label: 'Malleability', abbr: 'MA' },
  { key: 'pe', label: 'Potential Energy', abbr: 'PE' },
  { key: 'oq', label: 'Overall Quality', abbr: 'OQ' },
  { key: 'sr', label: 'Shock Resistance', abbr: 'SR' },
  { key: 'ut', label: 'Unit Toughness', abbr: 'UT' },
  { key: 'er', label: 'Entangle Resistance', abbr: 'ER' },
] as const

export type StatKey = typeof RESOURCE_STATS[number]['key']

// Resource types are now fetched from the /v1/resource-types API endpoint.
// See ui/src/features/resource-types/ for the ResourceType interface and API client.
// The resource_types table contains 1400+ types with stat ranges, categories, and groups.
