export const base64ToId = (str: string): number => {
  const decodedString = atob(str)
  const regex = /:(\d+)$/
  const matches = decodedString.match(regex)
  if (matches && matches.length > 1) {
    return parseInt(matches[1])
  } else {
    return 0
  }
}

export const idToBase64 = (id: number, prefix: string): string => {
  return btoa(`${prefix}:${id}`)
}

export const chinchiroRoomStatuses = new Map<string, string>([
  ['Opened', '公開中'],
  ['Finished', '終了']
])

export const convertToChinchiroRoomStatusName = (code: string): string | null =>
  chinchiroRoomStatuses.get(code) || null

export const playerAuthorities = new Map<string, string>([
  ['AuthorityPlayer', 'プレイヤー'],
  ['AuthorityAdmin', '管理者']
])
