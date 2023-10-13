import { User } from "@/types"

import { transformUser } from "@/lib/transform"

export async function getUser(
  resource: RequestInfo,
  init?: RequestInit
): Promise<User | null> {
  try {
    const res = await fetch(resource, init)

    if (!res.ok) {
      throw new Error(res.statusText)
    }

    const data = await res.json()

    const review = transformUser(data)

    return review
  } catch (error) {
    console.error(error)
    return null
  }
}