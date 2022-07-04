<script>
  import { createEventDispatcher } from 'svelte'
  import StandardTable from './variant/standard-table/index.svelte'
  import { SortOrder } from './utils'
  import { filterData, sortData, paginateData } from './hooks'

  const dispatch = createEventDispatcher()

  //   export let rowIconActions = true
  //   export let selectable = false
  //   export let i18n = {}
  //   export let selectedRowIds = []
  //   export let onRowSelect
  //   export let getRowIconActions = () => []
  //   export let onActionClick
  //   export let getRowClassName
  //
  //

  // core
  export let variant = 'standard'
  export let name
  export let colDefs = []
  export let data = []
  export let idField = 'id'
  export let renderCells = {}
  export let renderTypes = {}
  export let disabled = false

  // cosmetics - legacy
  export let maxHeight = -1
  export let fullWidth = true
  export let bgColorTable = '#FCFCFF'
  export let bgColorHead = '#FFFFFF'
  // more legacy
  export let selectable = false
  export let selectedRowIds = []
  export let rowIconActions = []

  // filters
  export let filters = {}
  export let filtersEnabled = true
  export let useServerFilters = false
  // sort
  export let sort = {}
  export let sortEnabled = true
  export let useServerSort = false
  export let sortByTypeFunctions = {}
  // paginate
  export let pagination = { page: 1, pageSize: 10 }
  export let paginationEnabled = true
  export let useServerPagination = false
  export let hasBoundaryRight = true
  export let pager = true // show pager?
  export let alignPager = 'center'

  // use server catch-all
  export let useServerAll = false

  // define preferences for server interaction
  let serverPagination = false
  let serverSort = false
  let serverFilters = false

  $: {
    serverPagination = useServerAll || useServerPagination
    serverSort = useServerAll || useServerSort
    serverFilters = useServerAll || useServerFilters
  }

  // filter
  let filteredData = data ? [...data] : []
  let filtersState = { ...filters }

  $: {
    filteredData = filterData(data, filtersEnabled, serverFilters, filtersState)
  }

  function updateFilters(key, filter) {
    if (!filtersEnabled) {
      return
    }
    filtersState = {
      ...filtersState,
      [key]: {
        ...(filtersState[key] || {}),
        ...filter,
      },
    }
    dispatch('filter', { name, filters: { ...filtersState } })
  }

  // sort
  let sortedData = [...filteredData]
  let sortState = { ...sort }

  $: {
    sortedData = sortData(
      filteredData,
      sortEnabled,
      serverSort,
      sortState,
      colDefs,
      sortByTypeFunctions
    )
  }

  function toggleSort(colId) {
    if (!sortEnabled) {
      return
    }
    let value = null
    if (sortState.sortColumn === colId) {
      value =
        sortState.sortOrder === SortOrder.asc ? SortOrder.desc : SortOrder.asc
    } else {
      value = SortOrder.desc
    }
    sortState = { sortColumn: colId, sortOrder: value }
    dispatch('sort', { name, colId, value })
  }

  // paginate
  let pageData = [...sortedData]
  let paginationState = { ...pagination }

  $: {
    pageData = paginateData(
      sortedData,
      paginationEnabled,
      serverPagination,
      paginationState
    )
  }

  function updatePagination(page) {
    paginationState = { page, pageSize: paginationState.pageSize }
    dispatch('paginate', { name, page, pageSize: paginationState.pageSize })
  }

  // render
  let tableVariants = {
    standard: StandardTable,
  }
  let renderComp = null

  $: {
    renderComp = tableVariants[variant]
  }

  let renderProps = {}
  $: {
    renderProps = {
      colDefs,
      data: pageData,
      idField,
      renderCells,
      renderTypes,
      disabled,
      maxHeight,
      fullWidth,
      bgColorTable,
      bgColorHead,
      selectable,
      selectedRowIds,
      filtersEnabled,
      filtersState,
      sortEnabled,
      sortState,
      paginationEnabled,
      paginationState,
      totalItems: sortedData.length,
      hasBoundaryRight,
      pager,
      alignPager,
      rowIconActions,
    }
  }

  function onFilterClick(e) {
    const colId = e.detail.colId
    console.log('onFilterClick: colId =', colId)
  }

  function onHeaderClick(e) {
    toggleSort(e.detail.colId)
  }

  function onPaginate(e) {
    updatePagination(e.detail.page)
  }

  function onAction(e) {
    const { type, value } = e.detail
    dispatch('action', { name, type, value })
  }
</script>

<div class="tui-table">
  {#if renderComp}
    <svelte:component
      this={renderComp}
      {...renderProps}
      on:filter={onFilterClick}
      on:header={onHeaderClick}
      on:paginate={onPaginate}
      on:action={onAction}
    />
  {:else}
    <div>Unknown table variant.</div>
  {/if}
</div>

<style>
  .tui-table {
    width: 100%;
  }

  :global(.green) {
    color: green;
  }
  :global(.red) {
    color: red;
  }
  :global(.link) {
    cursor: pointer;
    color: blue;
  }
  :global(.num) {
    float: right;
    text-align: right;
    color: blue;
  }
</style>
