import { valueSet } from './types'

export const formatDate = (str, showSeconds = true) => {
  if (str === '0001-01-01T00:00:00.000Z') {
    return ''
  }
  return showSeconds
    ? new Date(str)
        .toISOString()
        .replace('T', ' ')
        .replace('Z', '')
        .split('.')[0]
    : new Date(str).toISOString().replace('T', ' ').slice(0, 16)
}

export const addNumCommas = (value, round = -1, defaultValue = '') => {
  if (!valueSet(value)) {
    return defaultValue
  }
  if (round !== -1) {
    value = value.toFixed(round)
  }
  const parts = value.toString().split('.')
  parts[0] = parseInt(parts[0]).toLocaleString('en', { useGrouping: true })
  return parts.join('.')
}

export const formatNum = (value, defaultValue = '') => {
  return valueSet(value) ? addNumCommas(value) : defaultValue
}

export const formatSatoshi = (value, defaultValue = '') => {
  return valueSet(value) ? addNumCommas(value / 1e8, 8) : defaultValue
}

export const getTransactionHashUrl = (value, defaultValue = '') => {
  return valueSet(value) ? `https://whatsonchain.com/tx/${value}` : defaultValue
}

export const formatTransactionHash = (value, defaultValue = '') => {
  return valueSet(value)
    ? `${value.substring(0, 8)}...${value.substring(
        value.length - 8,
        value.length
      )}`
    : defaultValue
}

export const getWalletUrl = (value, defaultValue = '') => {
  return valueSet(value)
    ? `https://whatsonchain.com/address/${value}`
    : defaultValue
}

export const getBlockUrl = (value, defaultValue = '') => {
  return valueSet(value)
    ? `https://whatsonchain.com/block/${value}`
    : defaultValue
}

export const humanHash = (val) => {
  let unit = 'H/s'

  if (isNaN(val)) {
    return '-'
  }

  if (val >= 1e21) {
    val = val / 1e21
    unit = 'ZH/s'
  } else if (val >= 1e18) {
    val = val / 1e18
    unit = 'EH/s'
  } else if (val >= 1e15) {
    val = val / 1e15
    unit = 'PH/s'
  } else if (val >= 1e12) {
    val = val / 1e12
    unit = 'TH/s'
  } else if (val >= 1e9) {
    val = val / 1e9
    unit = 'GH/s'
  } else if (val >= 1e6) {
    val = val / 1e6
    unit = 'MH/s'
  } else if (val >= 1e3) {
    val = val / 1e3
    unit = 'kH/s'
  }

  return val.toFixed(2) + ' ' + unit
}

export const petaHash = (val) => {
  if (isNaN(val)) {
    return '-'
  }
  return (val / 1e15).toFixed(2) + ' PH/s'
}

export const dataSize = (val) => {
  let unit = 'B'

  if (isNaN(val)) {
    return '-'
  }

  if (val >= 1.1258999068426e15) {
    val = val / 1.1258999068426e15
    unit = 'PB'
  } else if (val >= 1099511627776) {
    val = val / 1099511627776
    unit = 'TB'
  } else if (val >= 1073741824) {
    val = val / 1073741824
    unit = 'GB'
  } else if (val >= 1048576) {
    val = val / 1048576
    unit = 'MB'
  } else if (val >= 1024) {
    val = val / 1024
    unit = 'kB'
  }

  return val.toFixed(2) + ' ' + unit
}

export const link = (href, text = null, className, external = true) => {
  const propObj = external
    ? { target: '_blank', rel: 'noopener noreferrer' }
    : {}
  let prefix = ''
  if (external) {
    try {
      const url = new URL(href)
      if (!url.protocol) {
        prefix = 'https://'
      }
    } catch (e) {
      prefix = 'https://'
    }
  }
  return `<a href='${prefix + href}'${
    external ? " target='_blank' rel='noopener noreferrer'" : ''
  }${className ? ` class='${className}'` : ''}>${text || href}</a>`
}
