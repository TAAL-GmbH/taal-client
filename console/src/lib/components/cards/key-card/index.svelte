<script>
  import Field from '../../field/index.svelte'
  import { dataSize, formatDate } from '../../../utils/format'

  export let testId = null

  export let key

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<div class="tui-key-card" {...optProps}>
  <Field label="Api key" value={key.api_key} copy={true} />
  <Field label="Data sent: {dataSize(key.dataBytes)}" layout="row" />

  <Field label="Address" value={key.address} copy={true} />
  <Field
    label="Date registered"
    value={formatDate(key.createdAt, false, false)}
  />
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
