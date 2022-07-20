<script>
  import { createEventDispatcher } from 'svelte'

  import { mediaSize } from '../../../../stores/index'
  import Button from '../../../button/index.svelte'
  import Checkbox from '../../../checkbox/index.svelte'
  import Icon from '../../../icon/index.svelte'
  import Pager from '../../../pager/index.svelte'
  import { getDisplay, SortOrder } from '../../utils'

  const dispatch = createEventDispatcher()

  export let name
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
  export let getRowIconActions
  export let getRenderProps
  // export let getRowClassName

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

  function onActionIcon(type, value) {
    dispatch('action', { name, type, value })
  }

  let gutter = 32
  let iconsW = 0

  $: {
    iconsW = 0
    let maxIconsW = 0
    data.forEach((item) => {
      const icons = getRowIconActions
        ? getRowIconActions(name, item, idField) || []
        : []
      const w = icons.length > 0 ? icons.length * 40 - 4 : 0
      if (w > maxIconsW) {
        maxIconsW = w
      }
    })
    iconsW = maxIconsW + 40
  }

  let gridGap = 0
  $: fixedWidth =
    iconsW +
    (selectable ? 32 : 0) +
    (colDefs.length > 0 ? (colDefs.length - 1) * gridGap : 0)
  let grid = ''

  $: {
    let gridItems = selectable ? ['32px'] : []
    if ($mediaSize === 'small') {
      gridItems.push(`calc(100% - ${selectable ? '80px' : '48px'})`)
      if (iconsW > 0) {
        gridItems.push(`48px`)
      }
    } else {
      colDefs.forEach((colDef) => {
        gridItems.push(
          colDef?.props?.width
            ? `calc(${colDef.props.width} - (${fixedWidth}px / ${colDefs.length}))`
            : `calc((100% - ${fixedWidth}px) / ${colDefs.length})`
        )
      })
      if (iconsW > 0) {
        gridItems.push(`${iconsW}px`)
      }
    }
    grid = gridItems.join(' ')
  }

  function getStyleProps(colDef) {
    if (colDef?.props?.style) {
      let str = ''
      Object.keys(colDef?.props?.style).forEach((key) => {
        str += `${key}: ${colDef.props.style[key]};`
      })
      return { style: str }
    }
    return {}
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
  style:--grid-local={grid}
  style:--grid-gap-local={gridGap + 'px'}
>
  <div
    class="table-container"
    class:fullWidth
    class:maxHeight={maxHeight !== -1 && !isNaN(maxHeight)}
  >
    <section
      class="table"
      class:small={$mediaSize === 'small'}
      class:selectable
      class:sortEnabled
      class:hasPager={paginationEnabled && pager}
    >
      {#if $mediaSize !== 'small'}
        <section class="thead">
          <div class="tr">
            {#if selectable}
              <div class="th" />
            {/if}
            {#each colDefs as colDef (colDef.id)}
              <div class="th" on:click={() => onHeaderClick(colDef.id)}>
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
              </div>
            {/each}
            {#if getRowIconActions}
              <div class="th" />
            {/if}
          </div>
        </section>
      {/if}
      <section class="tbody">
        {#each data as item (item[idField])}
          <div class="tr">
            {#if selectable}
              <div class="td">
                <Checkbox
                  name={item[idField]}
                  size="small"
                  checked={selectedRowIds.includes(item[idField])}
                  on:change={() => onRowSelect(item[idField])}
                />
              </div>
            {/if}
            {#if $mediaSize !== 'small'}
              {#each colDefs as colDef (colDef.id)}
                <div class="td">
                  {#if getDisplay(renderCells, renderTypes, colDef, idField, item).component}
                    <svelte:component
                      this={getDisplay(
                        renderCells,
                        renderTypes,
                        colDef,
                        idField,
                        item
                      ).component}
                      {...{
                        ...getDisplay(
                          renderCells,
                          renderTypes,
                          colDef,
                          idField,
                          item
                        ).props,
                        ...(getRenderProps
                          ? getRenderProps(name, colDef, idField, item)
                          : {}),
                      }}
                    />
                  {:else}
                    {getDisplay(renderCells, renderTypes, colDef, idField, item)
                      .value}
                  {/if}
                </div>
              {/each}
            {:else}
              <div class="td">
                <div class="inner-grid">
                  {#each colDefs as colDef (colDef.id)}
                    <div class="inner-grid-item" {...getStyleProps(colDef)}>
                      <div
                        class="inner-grid-item-label"
                        on:click={() => onHeaderClick(colDef.id)}
                      >
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
                      </div>
                      <div class="inner-grid-item-value">
                        {#if getDisplay(renderCells, renderTypes, colDef, idField, item).component}
                          <svelte:component
                            this={getDisplay(
                              renderCells,
                              renderTypes,
                              colDef,
                              idField,
                              item
                            ).component}
                            {...{
                              ...getDisplay(
                                renderCells,
                                renderTypes,
                                colDef,
                                idField,
                                item
                              ).props,
                              ...(getRenderProps
                                ? getRenderProps(name, colDef, idField, item)
                                : {}),
                            }}
                          />
                        {:else}
                          {getDisplay(
                            renderCells,
                            renderTypes,
                            colDef,
                            idField,
                            item
                          ).value}
                        {/if}
                      </div>
                    </div>
                  {/each}
                </div>
              </div>
            {/if}
            {#if getRowIconActions}
              <div class="td" class:column={$mediaSize === 'small'}>
                {#if !disabled}
                  <div class="table-cell-row">
                    {#each getRowIconActions(name, item, idField) || [] as actionItem (actionItem.icon)}
                      <div
                        class="action"
                        class:disabled={actionItem.disabled}
                        on:click={actionItem.disabled
                          ? null
                          : () => onActionIcon(actionItem.type, item)}
                      >
                        {#if actionItem.render === 'icon'}
                          <div class="ico">
                            <Icon name={actionItem.icon} size={18} />
                          </div>
                        {:else}
                          <div class="btn">
                            <Button
                              variant="ghost"
                              size="small"
                              disabled={actionItem.disabled}
                              icon={actionItem.icon}
                            />
                          </div>
                        {/if}
                      </div>
                    {/each}
                  </div>
                {/if}
              </div>
            {/if}
          </div>
        {/each}
      </section>
    </section>
  </div>
  {#if paginationEnabled && pager}
    <div class="table-pager">
      <Pager
        {totalItems}
        value={paginationState.page}
        pageSize={paginationState.pageSize}
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

    font-family: var(--font-family);
    box-sizing: var(--box-sizing);
  }

  .table-container {
    /* overflow-x: auto; */
    overflow-x: hidden;
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

  .table.maxHeight .thead .th {
    position: sticky;
    top: 0;
    box-shadow: 0 2px 1px -1px rgb(0 0 0 / 5%);
  }

  .table .tr {
    margin: 0;
    height: 60px;
    display: grid;
    grid-template-columns: var(--grid-local);
    gap: var(--grid-gap-local);
  }
  .table .tr:first-child .th:first-child,
  .table.small .tr:first-child .td:first-child {
    border-top-left-radius: 4px;
  }
  .table .tr:first-child .th:last-child,
  .table.small .tr:first-child .td:last-child {
    border-top-right-radius: 4px;
  }
  .table .tr:last-child .td:first-child {
    border-bottom-left-radius: 4px;
  }
  .table .tr:last-child .td:last-child {
    border-bottom-right-radius: 4px;
  }
  .table.hasPager .tr:last-child .td:first-child {
    border-bottom-left-radius: 0px;
  }
  .table.hasPager .tr:last-child .td:last-child {
    border-bottom-right-radius: 0px;
  }

  .table .th {
    text-align: left;
    white-space: nowrap;

    font-weight: 600;
    font-size: 13px;
    line-height: 16px;
    letter-spacing: 0.02em;

    overflow: hidden;
    display: flex;
    align-items: center;

    padding: 0px 32px;
    text-transform: uppercase;
    color: #232d7c;
    background-color: var(--bg-col-head-local);

    cursor: auto;
  }
  .table.selectable .th {
    padding: 5px;
  }
  .table.sortEnabled .th {
    cursor: pointer;
  }

  .table .th + .th {
    padding-left: 15px;
  }
  .table.selectable .th + .th {
    padding-left: 10px;
  }

  .table .td {
    font-weight: 400;
    font-size: 14px;
    line-height: 24px;
    letter-spacing: -0.01em;

    overflow: hidden;
    display: flex;
    align-items: center;

    color: #282933;

    padding: 0px 32px;
    border-top: 1px solid #efefef;
    border-bottom: 1px solid #fcfcff;
    background-color: var(--row-col-local);
  }
  .table.hasPager .tr:last-child .td {
    border-bottom: 1px solid #efefef;
  }
  .table.selectable .td {
    padding: 5px;
  }

  .table .td + .td {
    padding-left: 15px;
  }
  .table.selectable .td + .td {
    padding-left: 10px;
  }

  .table.small .tbody .td {
    border-top: 1px solid transparent;
    border-bottom: 1px solid transparent;
    padding: 0;
  }
  .table.small .tbody .tr {
    border-top: 1px solid #efefef;
    border-bottom: 1px solid #fcfcff;
    background-color: var(--row-col-local);
    height: inherit;
  }
  .table.small .tbody .tr:first-of-type {
    border-top: 1px solid transparent;
  }
  .inner-grid {
    display: grid;
    grid-template-columns: 50% 50%;
    row-gap: 8px;
    column-gap: 15px;
    /* display: flex;
    flex-wrap: wrap;
    gap: 8px; */
    width: 100%;
    height: calc(100% - 32px);
    padding: 16px 24px;
  }
  .inner-grid-item {
    display: flex;
    flex-direction: column;
    min-width: calc(50% - 4px);
  }
  .inner-grid-item-label {
    text-align: left;
    white-space: nowrap;

    font-weight: 600;
    font-size: 13px;
    line-height: 16px;
    letter-spacing: 0.02em;

    overflow: hidden;
    display: flex;
    align-items: center;

    text-transform: uppercase;
    color: #232d7c;

    cursor: auto;
  }
  .table.sortEnabled .inner-grid-item-label {
    cursor: pointer;
  }

  .table-cell-row {
    display: flex;
    align-items: center;
    flex-wrap: nowrap;
    gap: 4px;
  }
  .td.column {
    padding: 8px 16px 8px 0 !important;
  }
  .td.column .table-cell-row {
    flex-direction: column;
    justify-content: flex-start;
    height: 100%;
  }

  .header-icon {
    width: 18px;
    height: 18px;
    padding: 0 1px 0 3px;
  }
  .action {
    cursor: pointer;
    color: #232d7c;
  }
  .action .ico {
    width: 18px;
    height: 18px;
    margin: 9px;
  }
  .action .btn {
    width: 36px;
    height: 36px;
  }
  .action.disabled {
    cursor: auto;
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
