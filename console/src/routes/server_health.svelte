<script>
  import { onMount } from 'svelte'
  let srvStatusText
  let srvStatusClass

  let dbStatusText
  let dbStatusClass

  function setServerStatusRunning() {
    srvStatusText = 'Server running'
    srvStatusClass = 'button is-success'
  }

  function setServerStatusNotRunning() {
    srvStatusText = 'Server not running'
    srvStatusClass = 'button is-danger'
  }

  function setDBStatusRunning() {
    dbStatusText = 'DB running'
    dbStatusClass = 'button is-success'
  }

  function setDBStatusNotRunning() {
    dbStatusText = 'DB not running'
    dbStatusClass = 'button is-danger'
  }

  onMount(() => {
    const interval = setInterval(() => {
      const r = fetch(`${BASE_URL}/api/v1/health`)
        .then((res) => {
          if (!res.ok) {
            return res.text().then((text) => {
              throw new Error(text)
            })
          } else {
            setServerStatusRunning()
            setDBStatusRunning()
          }
        })
        .catch((err) => {
          try {
            const errJson = JSON.parse(err.message)
            setServerStatusRunning()
            setDBStatusNotRunning()            
          } catch (error) {
            setServerStatusNotRunning()
            setDBStatusNotRunning()
          }
        })
    }, 1000)

    return () => {
      clearInterval(interval)
    }
  })
</script>

<div class="field">
  <label for="serverStatus">Server status: </label>
  <button id="serverStatus" class={srvStatusClass}>{srvStatusText}</button>
</div>

<div class="field">
  <label for="dbStatus">Database status: </label>
  <button id="dbStatus" class={dbStatusClass}>{dbStatusText}</button>
</div>
