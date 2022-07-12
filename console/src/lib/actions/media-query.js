import { readable } from 'svelte/store'

export function query(query) {
  const mediaQueryList = window.matchMedia(query)

  const mediaStore = readable(mediaQueryList.matches, (set) => {
    const handleChange = () => set(mediaQueryList.matches)

    mediaQueryList.addEventListener('change', handleChange)

    return () => mediaQueryList.removeEventListener('change', handleChange)
  })

  return mediaStore
}
