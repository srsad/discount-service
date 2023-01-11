export enum SourceTypesEnum {
  STORE = 'store',
  PICKUP_POINT = 'pickup-point',
}

export type SourceTypes = keyof typeof SourceTypesEnum

export interface SourceRequest {
  name: string
  type: string
}

export interface Source {
  id: number
  name: string
  type: string
  created_at: string
  updated_at: string
}
