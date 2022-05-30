<script>
  import { onMount } from 'svelte'

  import HealthStatus from './health_status/health_status.svelte'

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

<div class="field">
  <label for="serverStatus">Server status: </label>
  <HealthStatus id="serverStatus" status={serverStatus} />
</div>
<div class="field">
  <label for="dbStatus">DB status: </label>
  <HealthStatus id="dbStatus" status={dbStatus} />
</div>
