<script>
  import { createEventDispatcher } from 'svelte'

  import Button from '../../button/index.svelte'
  import Field from '../../field/index.svelte'
  import { formatDate } from '../../../utils/format'

  const dispatch = createEventDispatcher()

  export let key

  function onAction() {
    dispatch('deactivate', { key })
  }
</script>

<div class="tui-key-card">
  <Field label="Api key" value={key.api_key} copy={true} />
  <Field label="Data sent: 20.5 MB" layout="row" />

  <Field label="Address" value={key.address} copy={true} />
  <Field
    label="Date registered"
    value={formatDate(key.createdAt, false, false)}
  />
  <Button variant="secondary" size="small" forceRed on:click={onAction}>
    Deactivate
  </Button>
</div>

<style>
  .tui-key-card {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);

    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 24px;

    width: 100%;
    padding: 24px;

    background: #fcfcff;
    border-radius: 4px;
  }
</style>
