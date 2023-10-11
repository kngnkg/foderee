import { Review } from "@/types"

import { env } from "@/env.mjs"

const transformReview = (apiReview: any): Review => {
  return {
    reviewId: apiReview.review_id,
    author: {
      userId: apiReview.author.user_id,
      displayId: apiReview.author.display_id,
      name: apiReview.author.name,
      avatarUrl: apiReview.author.avatar_url,
      bio: apiReview.bio ?? "", // デフォルト値
      followersCount: apiReview.author.followers_count,
      followingCount: apiReview.author.following_count,
      followed: apiReview.author.followed,
      following: apiReview.author.following,
      createdAt: apiReview.author.created_at
        ? new Date(apiReview.author.created_at)
        : new Date(),
      updatedAt: apiReview.author.updated_at
        ? new Date(apiReview.author.updated_at)
        : new Date(),
    },
    album: {
      albumId: apiReview.album.album_id,
      spotifyUri: apiReview.album.spotify_uri,
      spotifyUrl: apiReview.album.spotify_url,
      name: apiReview.album.name,
      diskNumber: apiReview.album.disk_number,
      artists: apiReview.album.artists.map((artist: any) => ({
        artistId: artist.artist_id,
        spotifyUri: artist.spotify_uri,
        name: artist.name,
        imageUrl: artist.image_url,
        genres: artist.genres,
      })),
      tracks: [], // デフォルト値
      coverUrl: apiReview.album.cover_url,
      releaseDate: new Date(apiReview.album.release_date),
      genres: apiReview.album.genres,
    },
    title: apiReview.title,
    body: "", // デフォルト値
    likesCount: apiReview.likesCount,
    liked: apiReview.liked,
    createdAt: new Date(apiReview.created_at),
    updatedAt: new Date(apiReview.updated_at),
  }
}

async function getReviews(): Promise<Review[] | null> {
  try {
    const res = await fetch(`${env.API_ROOT}/trends`)

    if (!res.ok) {
      throw new Error(res.statusText)
    }

    const data = await res.json()

    const reviews: Review[] = data.reviews.map((review: any) =>
      transformReview(review)
    )

    return reviews
  } catch (error) {
    console.error(error)
    return null
  }
}

export default async function IndexPage() {
  const reviews = await getReviews()

  if (!reviews) {
    return <p>Something went wrong.</p>
  }

  return (
    <>
      <section>
        <div className="">
          {reviews.map((review, idx) => {
            return (
              <div key={idx} className="m-2 pb-2 border-b">
                <p>タイトル: {review.title}</p>
                <p>アルバムアート: {review.album.coverUrl}</p>
                <p>ユーザー: {review.author.name}</p>
                <p>displayId: {review.author.displayId}</p>
                <p>userId: {review.author.userId}</p>
                <p>avatarUrl: {review.author.avatarUrl}</p>
                <p>likesCount: {review.likesCount}</p>
                <p>liked: {review.liked}</p>
                {review.album.artists.map((artist, idx) => {
                  return <p key={idx}>アーティスト: {artist.name}</p>
                })}
                <p>created_at: {review.createdAt.toDateString()}</p>
                {review.album.genres.map((genre, idx) => {
                  return <p key={idx}>ジャンル: {genre}</p>
                })}
              </div>
            )
          })}
        </div>
      </section>
    </>
  )
}
