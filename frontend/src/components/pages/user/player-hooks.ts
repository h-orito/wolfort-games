import {
  MyPlayerDocument,
  MyPlayerQuery,
  Player
} from '@/lib/generated/graphql'
import { useLazyQuery } from '@apollo/client'
import { atom, useAtom, useAtomValue } from 'jotai'
import { useEffect } from 'react'

// player
const myPlayerAtom = atom<Player | null>(null)

export const useMyPlayer = (): Player | null => {
  const [fetchMyPlayer] = useLazyQuery<MyPlayerQuery>(MyPlayerDocument)
  const [myPlayer, setMyPlayer] = useAtom(myPlayerAtom)
  useEffect(() => {
    const fetch = async () => {
      const { data } = await fetchMyPlayer()
      if (data?.myPlayer == null) return
      setMyPlayer(data.myPlayer as Player)
    }
    fetch()
    return () => setMyPlayer(null)
  }, [])
  return myPlayer
}
export const useMyPlayerValue = () => useAtomValue(myPlayerAtom)
export const isAdmin = (myPlayer: Player | null) => {
  return myPlayer && myPlayer.authorityCodes.includes('AuthorityAdmin')
}
