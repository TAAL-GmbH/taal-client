<script>
  import { onMount } from 'svelte'

  let config

  onMount(async () => {
    const r = await fetch(`${BASE_URL}/api/v1/config`)
    const data = await r.json()
    keys = data.keys
  })
</script>

ListenAddress string TaalUrl string TaalTimeOut time.Duration DBType string
DBFile string DBHost string DBPort int DBName string DBUsername string
DBPassword string

<div>Database settings</div>
<div class="control">
  <label class="radio">
    <input type="radio" name="answer" />
    Local
  </label>
  <label class="radio">
    <input type="radio" name="answer" />
    Remote
  </label>
</div>
