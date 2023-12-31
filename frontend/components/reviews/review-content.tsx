import Link from "next/link"
import { Review } from "@/types"

import { FollowButton } from "../follow-button"
import { TimeStamp } from "../timestamp"
import { UserAvatar } from "../user-avatar"
import { AlbumArt } from "./album-art"
import { Content } from "./content"

interface ReviewContentProps {
  review: Review
}

export const ReviewContent: React.FC<ReviewContentProps> = ({ review }) => {
  const pathToUser = `/${review.author.username}`

  return (
    <article className="flex flex-col gap-8">
      {/* アルバム情報 */}
      <section className="flex flex-col gap-4 pt-8">
        <div className="flex flex-col gap-2 items-center">
          <AlbumArt
            className="w-80 h-80 sm:w-80 sm:h-80"
            album={review.album}
          />
          <div>
            <p>アルバムタイトル・アーティスト名</p>
          </div>
          <div>
            <p>ジャンル</p>
          </div>
          <p>曲一覧(collupsible)</p>
        </div>
      </section>
      {/* タイトル */}
      <section className="my-4">
        <h1 className="scroll-m-20 text-3xl lg:text-5xl font-extrabold tracking-tight">
          {review.title}
        </h1>
      </section>
      {/* 投稿ユーザーの情報 */}
      <section className="flex gap-2 text-sm sm:text-md text-zinc-400 dark:text-zinc-400">
        <Link href={pathToUser}>
          <UserAvatar user={review.author} />
        </Link>
        <div className="flex flex-col">
          <Link href={pathToUser}>{review.author.displayName}</Link>
          <div className="flex gap-2 items-center">
            <TimeStamp date={review.createdAt} />
          </div>
        </div>
        <FollowButton user={review.author} variant="link" />
      </section>
      {/* 本文 */}
      <section>
        <Content data={review.content} />
      </section>
    </article>
  )
}
