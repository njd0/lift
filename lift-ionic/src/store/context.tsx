import { useContext } from "react"

import { AppContext } from "./AppProvider"

export function useApp() {
  const context = useContext(AppContext)
  if (context === undefined || context === null) {
    throw new Error("useApp must be used within an AppProvider")
  }
  return context
}
