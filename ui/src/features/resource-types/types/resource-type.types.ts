// Resource type matching the backend API
export interface ResourceType {
  resourceType: string
  resourceTypeName: string
  resourceCategory: string
  resourceGroup: string
  enterable: boolean
  maxTypes: number
  crMin: number
  crMax: number
  cdMin: number
  cdMax: number
  drMin: number
  drMax: number
  flMin: number
  flMax: number
  hrMin: number
  hrMax: number
  maMin: number
  maMax: number
  peMin: number
  peMax: number
  oqMin: number
  oqMax: number
  srMin: number
  srMax: number
  utMin: number
  utMax: number
  erMin: number
  erMax: number
  containerType: string
  inventoryType: string
  specificPlanet: number
}

export interface ResourceGroup {
  resourceGroup: string
  groupName: string
  groupLevel: number
  groupOrder: number
  containerType: string
}

export interface ResourceTypeListParams {
  page?: number
  rows?: number
  orderBy?: {
    field: string
    direction: 'ASC' | 'DESC'
  }
  resourceType?: string
  resourceTypeName?: string
  resourceCategory?: string
  resourceGroup?: string
  enterable?: boolean
  containerType?: string
}

export interface ResourceGroupListParams {
  page?: number
  rows?: number
  orderBy?: {
    field: string
    direction: 'ASC' | 'DESC'
  }
  resourceGroup?: string
  groupName?: string
  groupLevel?: number
  containerType?: string
}
