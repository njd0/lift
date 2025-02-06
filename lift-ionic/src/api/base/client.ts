import { handleFetchResponse } from "../base/errors"

export const API_HOST = "http://localhost:3333"

function getAbsoluteUrl(relativePath: string) {
  return new URL(relativePath, API_HOST).toString()
}

export type ApiConfig = {
  configured?: boolean
}

const Config: ApiConfig = {
  configured: false,
}

async function getAuthHeaders() {
  return {
    "X-Api-Key": "123", // todo add temp server auth
  }
}

export type StreamCallback<T> = (chunk: T) => void

export const BaseClient = {
  async patch(relativePath: string, init: RequestInit = {}) {
    const res = await fetch(getAbsoluteUrl(relativePath), {
      method: "PATCH",
      ...init,
      headers: {
        "Content-Type": "application/json",
        // ...(await getAuthHeaders()),
        // ...init.headers,
      },
    })

    return await handleFetchResponse(res)
  },
  async post(relativePath: string, init: RequestInit = {}) {
    const res = await fetch(getAbsoluteUrl(relativePath), {
      method: "POST",
      ...init,
      headers: {
        "Content-Type": "application/json",
        // ...(await getAuthHeaders()),
        // ...init.headers,
      },
    })

    return await handleFetchResponse(res)
  },
  async put() {},
  async package() {},
  async delete() {},
  async get(relativePath: string, init: RequestInit = {}) {
    const res = await fetch(getAbsoluteUrl(relativePath), {
      method: "GET",
      ...init,
      headers: {
        "Content-Type": "application/json",
        // ...(await getAuthHeaders()),
        // ...init.headers,
      },
    })

    return await handleFetchResponse(res)
  },
}
