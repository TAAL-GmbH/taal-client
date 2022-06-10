<script>
  import { onMount } from 'svelte'

  import HealthStatus from './health_status.svelte'

  let serverStatus = statusPending
  let dbStatus = statusPending

  const statusHealthy = 'healthy'
  const statusUnhealthy = 'unhealthy'
  const statusPending = 'pending'

  onMount(() => {
    const interval = setInterval(() => {
      const r = fetch(`${BASE_URL}/api/v1/health`)
        .then((res) => {
          if (!res.ok) {
            return res.text().then((text) => {
              throw new Error(text)
            })
          } else {
            serverStatus = statusHealthy
            dbStatus = statusHealthy
          }
        })
        .catch((err) => {
          try {
            const errJson = JSON.parse(err.message)
            serverStatus = statusHealthy
            dbStatus = statusUnhealthy
          } catch (error) {
            serverStatus = statusUnhealthy
            dbStatus = statusUnhealthy
          }
        })
    }, 1000)

    return () => {
      clearInterval(interval)
    }
  })
</script>

<div id="status">
  <div id="server" class="field">
    <label for="serverStatus">Server</label>
    <HealthStatus id="serverStatus" status={serverStatus} />
  </div>
  <div class="field">
    <label for="dbStatus">DB Connection</label>
    <HealthStatus id="dbStatus" status={dbStatus} />
  </div>
</div>

<style>
  #server {
    float: left;
    margin-right: 15px;
  }

  #status {
    margin-top: 20px;
  }
</style>
