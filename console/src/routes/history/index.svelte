<script>
  import { onMount } from 'svelte'
  import mime from 'mime'

  import ButtonSelect from '../../lib/components/button-select/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Table from '../../lib/components/table/index.svelte'

  import {
    downloadFileBlob,
    filterUnique,
    getColorFromDistinct,
  } from '../../lib/utils'
  import { success, failure } from '../../lib/utils/notifications'
  import * as api from '../../lib/api'
  import { getColDefs, getRangeItems } from './data'
  import i18n from '../../lib/i18n'

  $: t = $i18n.t
  const pageKey = 'page.history'

  let colDefs = []
  $: colDefs = getColDefs(t) || []

  let rangeItems = []
  $: rangeItems = getRangeItems(t) || []

  // injected by svelte-navigator
  export let location = null
  export let navigate = null

  let transactions = []
  let rangeValue = '720'
  let distinctKeys = []

  $: {
    distinctKeys = transactions.map((tx) => tx.api_key).filter(filterUnique)
  }

  function onRange(e) {
    getTransactions(e.detail.value)
  }

  function onAction(e) {
    download(e.detail.value)
  }

  const getRowIconActions = (tableName, item, idField = 'id') => {
    const icons = []
    if (item?.isHash) {
      icons.push({
        icon: 'hashtag',
        render: 'icon',
        disabled: true,
      })
    } else {
      icons.push({
        icon: 'download',
        type: 'download',
      })
    }
    if (item?.secret) {
      icons.push({
        icon: 'key',
        render: 'icon',
        disabled: true,
      })
    }
    return icons
  }

  const getRenderProps = (name, colDef, idField, item) => {
    let obj = {}
    if (colDef.id === 'api_key') {
      obj.statusColor = getColorFromDistinct(item.api_key, distinctKeys)
    }
    if (colDef.id === 'id') {
      obj.href = item.api_key.includes('testnet')
        ? `https://test.whatsonchain.com/tx/${item.id}`
        : `https://whatsonchain.com/tx/${item.id}`
    }
    return obj
  }

  function getTransactions(hours) {
    api.getTransactions(
      hours,
      (data) => {
        transactions = data.transactions
      },
      (error) =>
        failure(t('notifications.failure', { error }), { duration: 2000 })
    )
  }

  function download(tx) {
    let contentType = ''

    fetch(`${BASE_URL}/api/v1/transactions/${tx.id}`, {
      method: 'GET',
      headers: {
        Authorization: 'Bearer ' + tx.api_key,
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(t(`${pageKey}.errors.network-response`))
        }
        contentType = response.headers.get('content-type')
        return response.blob()
      })
      .then((blob) => {
        let filename = ''

        if (tx.filename === '') {
          const extension = mime.getExtension(contentType)
          const filenameSuffix = extension ? '.' + extension : ''

          filename = tx.id + filenameSuffix
        } else {
          filename = tx.filename
        }

        downloadFileBlob(filename, blob)
        success(t(`${pageKey}.notifications.download-success`, { filename }), {
          duration: 2000,
        })
      })
      .catch((error) =>
        failure(t('notifications.failure', { error }), { duration: 2000 })
      )
  }

  onMount(() => getTransactions(rangeValue))
</script>

<PageWithMenu>
  <div class="island">
    <Heading value={t(`${pageKey}.history-label`)} />
    <Spacer h={24} />
    <div class="sub-row">
      <ButtonSelect
        name="range"
        value={rangeValue}
        items={rangeItems}
        on:change={onRange}
      />
    </div>
    <Spacer h={24} />
    <Table
      name="history"
      variant="div"
      {colDefs}
      data={transactions}
      pagination={{
        page: 1,
        pageSize: 15,
      }}
      {getRenderProps}
      {getRowIconActions}
      on:action={onAction}
    />
  </div>
</PageWithMenu>

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .sub-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
</style>
