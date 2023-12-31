import { redisClient } from "@/lib/kvs"

import { setSpotifyClientAccessToken, spotifyClient } from "../spotify-client"

export default async function getAlbum(
  albumId: string
): Promise<SpotifyApi.SingleAlbumResponse> {
  try {
    await setSpotifyClientAccessToken(redisClient, spotifyClient)

    const resp = await spotifyClient.getAlbum(albumId)
    const data = resp.body

    // const resp = await fetch(`${env.MOCK_API_ROOT}/albums/${albumId}`, {})
    // const data = await resp.json()

    if (!data) {
      throw new Error("Failed to fetch album")
    }

    return data
  } catch (e) {
    console.error(e)
    throw new Error("Failed to fetch album")
  }
}
