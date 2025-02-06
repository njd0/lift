export type NinjaValidationErrorType = {
  loc: string[]
  msg: string
  type: string
}

function hasMessage(data: unknown): data is { message: string } {
  return (
    typeof data === "object" &&
    data !== null &&
    "message" in data &&
    typeof data.message === "string"
  )
}

export class ApiError extends Error {
  public status: number
  public data: unknown

  constructor(status: number, data: unknown, defaultMessage: string) {
    super(hasMessage(data) ? data.message : defaultMessage)
    this.name = new.target.name // get the name of the specific subclass
    this.status = status
    this.data = data
  }
}
class BadRequestError extends ApiError {
  constructor(data: unknown) {
    super(400, data, "Bad Request")
  }
}

class UnauthorizedError extends ApiError {
  constructor(data: unknown) {
    super(401, data, "Unauthorized")
  }
}

class ForbiddenError extends ApiError {
  constructor(data: unknown) {
    super(403, data, "Forbidden")
  }
}

class NotFoundError extends ApiError {
  constructor(data: unknown) {
    super(404, data, "Not Found")
  }
}

export class ValidationError extends ApiError {
  constructor(data: NinjaValidationErrorType[]) {
    super(422, data, "Validation Error")
  }
}

export async function throwForStatus(res: Response) {
  if (!res.ok) {
    const data = await res.json()
    switch (res.status) {
      case 400:
        throw new BadRequestError(data)
      case 401:
        throw new UnauthorizedError(data)
      case 403:
        throw new ForbiddenError(data)
      case 404:
        throw new NotFoundError(data)
      case 422:
        throw new ValidationError(data)
      default:
        throw new ApiError(res.status, data, "Api Error")
    }
  }
}

export async function handleFetchResponse(res: Response) {
  if (res.ok) {
    return await res.json()
  }

  await throwForStatus(res)
}
