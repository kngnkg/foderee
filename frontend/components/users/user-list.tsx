import { Suspense } from "react"
import Link from "next/link"
import { User } from "@/types"

import { useUsers } from "@/hooks/users/use-users"
import { FollowButton } from "@/components/follow-button"
import { UserAvatar } from "@/components/user-avatar"

import { Button } from "../ui/button"
import { UserListSkeleton } from "./user-list-skeleton"
import { UserSkeleton } from "./user-skeleton"

interface UserListProps {
  endpoint: string
}

export const UserList: React.FC<UserListProps> = ({ endpoint }) => {
  const { data, error, isLoading, isValidating, loadMore } = useUsers({
    endpoint: endpoint,
  })

  if (error) {
    console.error(error)
    return <p>Something went wrong.</p>
  }

  return (
    <div className="flex flex-col gap-4">
      {data ? (
        <>
          {data.map((userWP, idx) => (
            <>
              {userWP.users.map((user: User) => (
                <div key={idx} className="flex items-center justify-between">
                  <Suspense fallback={<UserSkeleton />}>
                    <div className="flex gap-2 items-center text-sm sm:text-md text-zinc-400 dark:text-zinc-400">
                      <Link href={`/${user.username}`}>
                        <UserAvatar user={user} />
                      </Link>
                      <div className="flex flex-col">
                        <Link href={`/${user.username}`}>
                          {user.displayName}
                        </Link>
                      </div>
                    </div>
                    <FollowButton user={user} />
                  </Suspense>
                </div>
              ))}
            </>
          ))}
        </>
      ) : (
        <>{isLoading ? <UserListSkeleton count={10} /> : <p>No Content.</p>}</>
      )}
      <div className="mb-20 flex flex-col items-center">
        <Button variant="ghost" size="lg" onClick={() => loadMore()}>
          もっと見る
        </Button>
      </div>
    </div>
  )
}
