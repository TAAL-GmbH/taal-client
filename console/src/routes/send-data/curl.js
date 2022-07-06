export function getCurlCommand(key, type, tag, mode, secret, data, url, file) {
  let curl = 'curl \\\n  -X POST \\\n'
  if (key) curl += `  -H 'Authorization: Bearer ${key}' \\\n`
  if (type) curl += `  -H 'Content-Type: ${type}' \\\n`
  if (tag) curl += `  -H 'X-Tag: ${tag}' \\\n`
  switch (mode) {
    case 'hash':
      curl += `  -H 'X-Mode: ${mode}' \\\n`
      break

    case 'encrypt':
      curl += `  -H 'X-Mode: ${mode}' \\\n`
      curl += `  -H 'X-Key: ${secret}' \\\n`
      break
  }

  if (file) {
    curl += `  --data-binary @'${file.name}' \\\n`
  } else if (data) {
    curl += `  -d '${data}' \\\n`
  }
  curl += `${url}/api/v1/transactions`

  return curl
}
