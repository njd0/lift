export type OptionalPK = number | string | null | undefined
export type LazyPK = number | undefined

export type BaseQueryHookProps<T = unknown> = T & {
  disabled?: boolean
}

export type ListResponse<T> = {
  count: number
  items: T
}
