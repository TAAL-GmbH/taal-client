<script>
  import { onMount } from 'svelte'
  import ButtonSelect from '../../lib/components/button-select/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Table from '../../lib/components/table/index.svelte'
  import Spinner from '../../lib/components/spinner/index.svelte'

  import { spinCount } from '../../lib/stores'
  import * as api from '../../lib/api'
  import { colDefs } from './data'

  let transactions = []
  let rangeValue = '720'

  function onRange(e) {
    rangeValue = e.detail.value
    getTransactions(rangeValue)
  }

  function onAction(e) {
    const { name, type, value } = e.detail
    console.log('onAction: type = ', type, ' value = ', value)
  }

  function getTransactions(hours) {
    api.getTransactions(
      hours,
      (data) => {
        transactions = data.transactions
      },
      (error) => {
        addNotification({
          text: `Error: ${errJson.error}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })
      }
    )
  }

  onMount(() => getTransactions(rangeValue))
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="History" />
    <Spacer h={24} />
    <div class="sub-row">
      <ButtonSelect
        name="range"
        value={rangeValue}
        items={[
          { label: '24H', value: '24' },
          { label: '1W', value: '168' },
          { label: '1M', value: '720' },
          { label: 'All', value: '' },
        ]}
        on:change={onRange}
      />
    </div>
    <Spacer h={24} />
    <Table
      {colDefs}
      data={transactions}
      pagination={{
        page: 1,
        pageSize: 15,
      }}
      rowIconActions={[
        { icon: 'download', name: 'Download', type: 'download' },
      ]}
      on:action={onAction}
    />
  </div>
</PageWithMenu>

{#if $spinCount > 0}
  <Spinner />
{/if}

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
