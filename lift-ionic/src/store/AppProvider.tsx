"use client"
import { createContext } from "react"
import { useImmer } from "use-immer"

import type { PropsWithChildren } from "react"

interface AppState {
  initialized: boolean
}

const initialState: AppState = {
  initialized: false,
}

export const AppContext = createContext<AppState>(initialState)

export default function AppProvider({ children }: PropsWithChildren) {
  const [state, setState] = useImmer<AppState>(initialState)

  return (
    <AppContext.Provider
      value={{
        ...state,
      }}>
      {children}
    </AppContext.Provider>
  )
}
