import { writable } from 'svelte/store'

// header
export const menuLinks = writable([])
export const menuActions = writable([])

// content
export const pageContentOffsetX = writable(0)

// spinner
export const spinCount = writable(0)
