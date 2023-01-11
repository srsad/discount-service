import type { AxiosResponse } from 'axios'

import apiClient from '@/apiClient'
import type { Response } from '@/types/common'
import type { Source, SourceRequest } from '@/types/sourse'

export default {
  // TODO: упростить - AxiosResponse<Response<any>>
  create(params: SourceRequest): Promise<AxiosResponse<Response>> {
    return apiClient.post('/source/', params)
  },

  getAll(): Promise<AxiosResponse<Response<Source[]>>> {
    return apiClient.get('/source/')
  },

  getById(sourceId: number): Promise<AxiosResponse<Response<Source>>> {
    return apiClient.get('/source/' + sourceId)
  },

  remove(sourseId: number): Promise<AxiosResponse<Response>> {
    return apiClient.delete('/source/' + sourseId)
  },

  update(params: Source): Promise<AxiosResponse<Response<Source>>> {
    return apiClient.put('/source/' + params.id, params)
  },
}
