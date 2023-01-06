/**
 * Общий интерфейс ответа
 */
export interface Response<T = null> {
  status: string
  message: string
  data: T
}
