import {
  ChinchiroGame,
  ChinchiroGameStatus,
  ChinchiroRoom,
  ChinchiroRoomParticipant,
  MyChinchiroRoomParticipantDocument,
  MyChinchiroRoomParticipantQuery,
  MyChinchiroRoomParticipantQueryVariables,
  QChinchiroGamesDocument,
  QChinchiroGamesQuery
} from '@/lib/generated/graphql'
import { useLazyQuery } from '@apollo/client'
import { atom, useAtom, useAtomValue, useSetAtom } from 'jotai'
import { useEffect } from 'react'

// room
const chinchiroRoomAtom = atom<ChinchiroRoom | null>(null)
export const useChinchiroRoom = (
  chinchiroRoom: ChinchiroRoom
): ChinchiroRoom => {
  const setChinchiroRoom = useSetAtom(chinchiroRoomAtom)
  setChinchiroRoom(chinchiroRoom)
  useEffect(() => {
    return () => setChinchiroRoom(null)
  }, [])
  return chinchiroRoom
}
export const useChinchiroRoomValue = () => useAtomValue(chinchiroRoomAtom)

// my room participant
const myChinchiroRoomParticipantAtom = atom<ChinchiroRoomParticipant | null>(
  null
)
export const useMyChinchiroRoomParticipantInit = (
  roomId: string
): ChinchiroRoomParticipant | null => {
  const [myself, refetchMyself] = useMyChinchiroRoomParticipant(roomId)
  const setMyselfAtom = useSetAtom(myChinchiroRoomParticipantAtom)
  useEffect(() => {
    refetchMyself()
    return () => setMyselfAtom(null)
  }, [])
  return myself
}
export const useMyChinchiroRoomParticipant = (
  roomId: string
): [myself: ChinchiroRoomParticipant | null, refetchMyself: () => void] => {
  const [fetchMyself] = useLazyQuery<MyChinchiroRoomParticipantQuery>(
    MyChinchiroRoomParticipantDocument
  )
  const [myself, setMyselfAtom] = useAtom(myChinchiroRoomParticipantAtom)
  const fetch = async () => {
    const { data } = await fetchMyself({
      variables: { roomId } as MyChinchiroRoomParticipantQueryVariables
    })
    setMyselfAtom(
      (data?.myChinchiroRoomParticipant as ChinchiroRoomParticipant) ?? null
    )
  }
  return [myself, fetch]
}

// room master
export const useIsRoomMaster = (roomId: string): boolean => {
  const room = useChinchiroRoomValue()
  const myself = useMyChinchiroRoomParticipantValue()
  return (
    room?.roomMasters?.some((rm) => rm.player.id === myself?.player.id) ?? false
  )
}

export const useMyChinchiroRoomParticipantValue = () =>
  useAtomValue(myChinchiroRoomParticipantAtom)

// games
const chinchiroGamesAtom = atom<Array<ChinchiroGame>>([])
export const useChinchiroGamesInit = (roomId: string): Array<ChinchiroGame> => {
  const [games, fetchGames] = useChinchiroGames(roomId)
  const setAtom = useSetAtom(chinchiroGamesAtom)
  useEffect(() => {
    fetchGames()
    return () => setAtom([])
  }, [])
  return games
}
export const useChinchiroGames = (
  roomId: string
): [games: Array<ChinchiroGame>, fetchGames: () => void] => {
  const [fetchGames] = useLazyQuery<QChinchiroGamesQuery>(
    QChinchiroGamesDocument
  )
  const [games, setGames] = useAtom(chinchiroGamesAtom)
  const fetch = async (): Promise<void> => {
    const { data } = await fetchGames({
      variables: {
        query: {
          roomId
        }
      }
    })
    setGames(data?.chinchiroGames as Array<ChinchiroGame>)
  }
  return [games, fetch]
}
export const useChinchiroGamesValue = () => useAtomValue(chinchiroGamesAtom)
export const useProgressChinchiroGameValue = (): ChinchiroGame | null => {
  const games = useChinchiroGamesValue()
  return (
    games.find((game) => game.status === ChinchiroGameStatus.Progress) ?? null
  )
}
