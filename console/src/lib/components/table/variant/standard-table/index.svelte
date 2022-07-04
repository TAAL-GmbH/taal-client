<script>
  import { createEventDispatcher } from 'svelte'
  import Checkbox from '../../../checkbox/index.svelte'
  import Icon from '../../../icon/index.svelte'
  import Pager from '../../../pager/index.svelte'
  import { getDisplayProps, SortOrder } from '../../utils'

  const dispatch = createEventDispatcher()

  export let colDefs = []
  export let data = []
  export let idField
  export let renderCells = {}
  export let renderTypes = {}
  export let disabled = false
  export let maxHeight = -1
  export let fullWidth
  export let bgColorTable
  export let bgColorHead
  export let selectable
  export let selectedRowIds = []
  export let filtersEnabled
  export let filtersState
  export let sortEnabled
  export let sortState
  export let paginationEnabled
  export let paginationState
  export let totalItems = -1
  export let hasBoundaryRight = true
  export let pager = true
  export let alignPager = 'center'
  export let rowIconActions = []

  function onFilterClick(colId) {
    dispatch('filter', { colId })
  }

  function onHeaderClick(colId) {
    dispatch('header', { colId })
  }

  function onRowSelect(id) {
    dispatch('select', { id })
  }

  function onPage(e) {
    dispatch('paginate', { page: e.detail.value })
  }
</script>

<div
  class="table-and-tools"
  style:--max-height-local={maxHeight + 'px'}
  style:--bg-col-head-local={bgColorHead
    ? bgColorHead
    : bgColorTable
    ? bgColorTable
    : '#ffffff'}
  style:--row-col-local={bgColorTable ? bgColorTable : 'none'}
  style:--align-pager-local={alignPager}
>
  <div
    class="table-container"
    class:fullWidth
    class:maxHeight={maxHeight !== -1 && !isNaN(maxHeight)}
  >
    <table
      cellpadding={0}
      cellspacing={0}
      class="table"
      class:selectable
      class:sortEnabled
      class:hasPager={paginationEnabled && pager}
    >
      <thead>
        <tr>
          {#if selectable}
            <th />
          {/if}
          {#each colDefs as colDef (colDef.id)}
            <th on:click={() => onHeaderClick(colDef.id)}>
              <div class="table-cell-row">
                {colDef.name}
                {#if sortEnabled && sortState.sortColumn === colDef.id}
                  <div class="header-icon">
                    <Icon
                      name={sortState.sortOrder === SortOrder.asc
                        ? 'chevron-up'
                        : 'chevron-down'}
                      size={18}
                    />
                  </div>
                {/if}
                {#if filtersEnabled && filtersState[colDef.id]}
                  <div class="header-icon">
                    <Icon
                      name="filters"
                      size={18}
                      on:click={() => onFilterClick(colDef.id)}
                    />
                  </div>
                {/if}
              </div>
            </th>
          {/each}
          {#if rowIconActions?.length > 0}
            <th />
          {/if}
        </tr>
      </thead>
      <tbody>
        {#each data as item (item[idField])}
          <tr>
            {#if selectable}
              <td>
                <Checkbox
                  name={item[idField]}
                  size="small"
                  checked={selectedRowIds.includes(item[idField])}
                  on:change={() => onRowSelect(item[idField])}
                />
              </td>
            {/if}
            {#each colDefs as colDef (colDef.id)}
              <td>
                {#if getDisplayProps(renderCells, renderTypes, colDef, idField, item).component}
                  <svelte:component
                    this={getDisplayProps(
                      renderCells,
                      renderTypes,
                      colDef,
                      idField,
                      item
                    ).component}
                    {...getDisplayProps(
                      renderCells,
                      renderTypes,
                      colDef,
                      idField,
                      item
                    ).props}
                  />
                {:else}
                  {getDisplayProps(
                    renderCells,
                    renderTypes,
                    colDef,
                    idField,
                    item
                  ).value}
                {/if}
              </td>
            {/each}
            {#if rowIconActions?.length > 0}
              <td>row icons</td>
            {/if}
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
  {#if paginationEnabled && pager}
    <div class="table-pager">
      <Pager
        {totalItems}
        {...paginationState}
        {hasBoundaryRight}
        on:change={onPage}
      />
    </div>
  {/if}
</div>

<style>
  .table-and-tools {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: stretch;
  }

  .table-container {
    overflow-x: auto;
  }
  .table-container.fullWidth {
    width: 100%;
  }
  .table-container.maxHeight {
    overflow-y: auto;
    max-height: var(--max-height-local);
  }

  .table {
    font-size: 14px;
    white-space: nowrap;
  }
  .table-container.fullWidth .table {
    width: 100%;
  }

  .table.maxHeight thead th {
    position: sticky;
    top: 0;
    box-shadow: 0 2px 1px -1px rgb(0 0 0 / 5%);
  }

  .table tr {
    margin: 0;
    height: 60px;
  }
  .table tr:first-child th:first-child {
    border-top-left-radius: 4px;
  }
  .table tr:first-child th:last-child {
    border-top-right-radius: 4px;
  }
  .table tr:last-child td:first-child {
    border-bottom-left-radius: 4px;
  }
  .table tr:last-child td:last-child {
    border-bottom-right-radius: 4px;
  }
  .table.hasPager tr:last-child td:first-child {
    border-bottom-left-radius: 0px;
  }
  .table.hasPager tr:last-child td:last-child {
    border-bottom-right-radius: 0px;
  }

  .table th {
    text-align: left;
    white-space: nowrap;

    font-weight: 600;
    font-size: 13px;
    line-height: 16px;
    letter-spacing: 0.02em;

    padding: 0px 32px;
    text-transform: uppercase;
    color: #232d7c;
    background-color: var(--bg-col-head-local);

    cursor: auto;
  }
  .table.selectable th {
    padding: 5px;
  }
  .table.sortEnabled th {
    cursor: pointer;
  }

  .table th + th {
    padding-left: 15px;
  }
  .table.selectable th + th {
    padding-left: 10;
  }

  .table td {
    font-weight: 400;
    font-size: 14px;
    line-height: 24px;
    letter-spacing: -0.01em;

    color: #282933;

    padding: 0px 32px;
    border-top: 1px solid #efefef;
    border-bottom: 1px solid #fcfcff;
    background-color: var(--row-col-local);
  }
  .table.hasPager tr:last-child td {
    border-bottom: 1px solid #efefef;
  }
  .table.selectable td {
    padding: 5px;
  }

  .table td input {
    vertical-align: middle;
    margin-top: -2px;
  }

  .table td + td {
    padding-left: 15px;
  }
  .table.selectable td + td {
    padding-left: 10px;
  }

  .table tr.selected {
    background: lightgreen;
  }
  .table tr.locked {
    background: oldlace;
  }
  .table tr.lockedSelf {
    background: lightcyan;
  }

  .table-cell-row {
    display: flex;
    align-items: center;
    flex-wrap: nowrap;
  }
  .header-icon {
    width: 18px;
    height: 18px;
    padding: 0 1px 0 3px;
  }
  .table-pager {
    width: 100%;
    height: 68px;
    display: flex;
    align-items: center;
    justify-content: var(--align-pager-local);

    background: #ffffff;
    border-radius: 0px 0px 4px 4px;
  }
</style>
