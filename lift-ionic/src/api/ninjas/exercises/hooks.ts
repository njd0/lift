import { useQuery } from "@tanstack/react-query"

import { ApiError } from "../../base/errors"
import * as exercisesApi from "../exercises/api"

import type { BaseQueryHookProps } from "../../base/types"

type SuccessType = {}
export function useGetAllExercises({
  params,
  disabled,
}: BaseQueryHookProps<{
  params: ExercisesGetParams
}>) {
  return useQuery<SuccessType, ApiError>({
    queryKey: ["Exercises"],
    queryFn: async () => await exercisesApi.getAllExercises(params),
    enabled: !disabled,
  })
}
