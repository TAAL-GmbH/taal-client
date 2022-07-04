<script>
  import ButtonSelect from '../../lib/components/button-select/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Table from '../../lib/components/table/index.svelte'

  let filtersOpen = false

  function onRange(e) {
    console.log('onRange: detail = ', e.detail)
  }

  function onFilters(e) {
    console.log('onFilters: detail = ', e.detail)
    filtersOpen = !filtersOpen
  }

  const tableColDefs = [
    {
      id: 'created_at',
      name: 'Created at (UTC)',
      type: 'dateStr',
    },
    {
      id: 'id',
      name: 'ID',
      type: 'transactionhash',
    },
    {
      id: 'api_key',
      name: 'API key used',
      type: 'apiKey',
    },
    {
      id: 'data_bytes',
      name: 'Data size',
      type: 'dataSize',
    },
    {
      id: 'filename',
      name: 'Filename',
      type: 'string',
    },
  ]

  const tableData = [
    {
      id: '111119bc38c0648706ef2de8d6668a7cf90d301ee5891c17530e9db96205dbc0',
      api_key: 'testnet_0b85b7619de806c0aae9075cbb0266cf',
      data_bytes: 213,
      created_at: '2022-06-22T08:46:36.259Z',
      filename: '',
      secret: '123',
    },
    {
      id: '222219bc38c0648706ef2de8d6668a7cf90d301ee5891c17530e9db96205dbc0',
      api_key: 'testnet_0b85b7619de806c0aae9075cbb0266cf',
      data_bytes: 123,
      created_at: '2022-06-19T08:46:36.259Z',
      filename: 'names.txt',
      secret: '456',
    },
    {
      id: '333319bc38c0648706ef2de8d6668a7cf90d301ee5891c17530e9db96205dbc0',
      api_key: 'testnet_0b85b7619de806c0aae9075cbb0266cf',
      data_bytes: 456,
      created_at: '2022-06-26T08:46:36.259Z',
      filename: '',
      secret: '789',
    },
    {
      id: '444419bc38c0648706ef2de8d6668a7cf90d301ee5891c17530e9db96205dbc0',
      api_key: 'testnet_0b85b7619de806c0aae9075cbb0266cf',
      data_bytes: 2233,
      created_at: '2022-05-24T08:46:36.259Z',
      filename: 'todo.doc',
      secret: '4545454',
    },
    {
      id: '555519bc38c0648706ef2de8d6668a7cf90d301ee5891c17530e9db96205dbc0',
      api_key: 'testnet_0b85b7619de806c0aae9075cbb0266cf',
      data_bytes: 34,
      created_at: '2022-06-30T08:46:36.259Z',
      filename: '',
      secret: 'abc',
    },
  ]

  function onAction(e) {
    const { name, type, value } = e.detail
    console.log('onAction: type = ', type, ' value = ', value)
  }
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="History" />
    <Spacer h={24} />
    <div class="sub-row">
      <ButtonSelect
        name="range"
        items={[
          { label: '24H', value: '24H' },
          { label: '1W', value: '1W' },
          { label: '1M', value: '1M' },
          { label: 'All', value: 'All' },
        ]}
        on:change={onRange}
      />
      <div style="display: flex; gap: 8px;">
        <ButtonSelect
          items={[{ label: 'Filters', value: 0, icon: 'filters' }]}
          focusRect={false}
          value={filtersOpen ? 0 : -1}
          on:change={onFilters}
        />
      </div>
    </div>
    <Spacer h={24} />
    <Table
      colDefs={tableColDefs}
      data={tableData}
      pagination={{
        page: 1,
        pageSize: 15,
      }}
      filters={{
        age: {
          filterType: 'anyOf',
          selectedValues: [],
        },
      }}
      rowIconActions={[
        { icon: 'download', name: 'Download', type: 'download' },
      ]}
      on:action={onAction}
    />
  </div>
</PageWithMenu>

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
    padding-top: 40px;
    margin-bottom: 100px;
  }

  .sub-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
</style>
