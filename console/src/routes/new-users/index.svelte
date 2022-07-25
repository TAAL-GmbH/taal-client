<script context="module">
  let firstLoad = true
</script>

<script>
  import { onMount } from 'svelte'
  import { useNavigate } from 'svelte-navigator'

  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageBasic from '../../lib/components/page/template/basic/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'

  import * as api from '../../lib/api'
  import i18n from '../../lib/i18n'

  $: t = $i18n.t
  const pageKey = 'page.new-users'

  const navigate = useNavigate()

  function onKey() {
    navigate('/register-key')
  }

  function onNoKey() {
    window.open('https://console.taal.com', '_blank')
  }

  onMount(async () => {
    const result = await api.getApiKeys()

    if (firstLoad && result?.keys && result?.keys.length > 0) {
      firstLoad = false
      navigate('/send-data')
    }
  })
</script>

<PageBasic>
  <div class="island">
    <Heading value={t(`${pageKey}.new-users-label`)} />
    <Spacer h={64} />
    <Row gap={16}>
      <Button variant="secondary" size="large" width={236} on:click={onNoKey}>
        {t(`${pageKey}.answer-no-label`)}
      </Button>
      <Button width={236} size="large" on:click={onKey}>
        {t(`${pageKey}.answer-yes-label`)}
      </Button>
    </Row>
  </div>
</PageBasic>

<style>
  .island {
    text-align: center;
  }
</style>
