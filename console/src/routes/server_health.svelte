<script>
  import { onMount } from 'svelte'
  let serverText
  let serverClass

  function setServerRunning() {
    serverText = 'Server running'
    serverClass = 'button is-success'
  }

  function setServerNotRunning() {
    serverText = 'Server not running'
    serverClass = 'button is-danger'
  }

  onMount(() => {
    const interval = setInterval(() => {
      const r = fetch(`${BASE_URL}/api/v1/health`)
        .then((res) => {
          if (!res.ok) {
            setServerNotRunning()
          } else {
            setServerRunning()
          }
        })
        .catch((err) => {
          setServerNotRunning()
        })
    }, 1000)

    return () => {
      clearInterval(interval)
    }
  })

</script>

<div class="fi">
  <label for="serverStatus">Server status: </label>
  <button id="serverStatus" class={serverClass}>{serverText}</button>
</div>
