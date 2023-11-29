import { NextRequest, NextResponse } from "next/server"
import listFollows from "@/service/follow/list-follows"

import { getServerSession } from "@/lib/session"
import {
  UserRouteContext,
  userRouteContextSchema,
} from "@/lib/validations/user"
import { errBadRequest, errInternal, errUnauthorized } from "@/app/api/response"

export async function GET(request: NextRequest, context: UserRouteContext) {
  try {
    const { params } = userRouteContextSchema.parse(context)

    const session = await getServerSession()
    if (!session || !session.idToken) {
      return errUnauthorized("unauthorized")
    }

    const resp = await listFollows(session.idToken, [params.username])
    if (!resp) {
      return errInternal("internal error")
    }

    return NextResponse.json({
      username: resp[0].getUsername(),
      immutableId: resp[0].getImmutableid(),
      displayName: resp[0].getDisplayname(),
      isFollowing: resp[0].getIsfollowing(),
    })
  } catch (e) {
    console.error(e)
    return errInternal("internal error")
  }
}

// POST
