import { BaseClient } from "../../base/client"

const baseRoute = "/v1/exercises"
type Typee = {}
export async function getAllExercises(
  params: ExercisesGetParams
): Promise<Typee> {
  const qp = new URLSearchParams(params)
  return await BaseClient.get(`${baseRoute}?${qp.toString()}`)
}
