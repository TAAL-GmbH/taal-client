<script>
  import Field from '../../field/index.svelte'
  import { dataSize, formatDate } from '../../../utils/format'
  import i18n from '../../../i18n'

  $: t = $i18n.t
  const tKey = 'comp.key-card'

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
  <Field label={t(`${tKey}.api-key-label`)} value={key.api_key} copy={true} />
  <Field
    label={t(`${tKey}.data-sent-label`, { size: dataSize(key.dataBytes) })}
    layout="row"
  />

  <Field label={t(`${tKey}.address-label`)} value={key.address} copy={true} />
  <Field
    label={t(`${tKey}.date-registered`)}
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
