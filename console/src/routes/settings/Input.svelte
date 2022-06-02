<script>
  import { beforeUpdate } from 'svelte'

  // FontAwesome icon...
  import Fa from 'svelte-fa'
  import { faXmark, faCheck } from '@fortawesome/free-solid-svg-icons'

  export let label
  export let description
  export let placeholder
  export let value
  export let updateFn

  let originalValue
  let currentValue

  beforeUpdate(() => {
    if (originalValue === undefined && currentValue === undefined) {
      originalValue = value
      currentValue = value
    }
  })

  function handleReset() {
    currentValue = originalValue
  }

  function handleSave() {
    originalValue = currentValue
    updateFn(currentValue)
  }
</script>

<div class="field is-horizontal">
  <div class="field-label is-normal has-text-left">
    <label for={label} class="label">{label}:</label>
  </div>
  <div class="field-body">
    <div class="field">
      <p class="control has-icons-right">
        <input
          id={label}
          class="input"
          type="text"
          {placeholder}
          bind:value={currentValue}
        />
        {#if currentValue !== originalValue}
          <span
            class="icon is-right"
            style="pointer-events: all; cursor: pointer"
          >
            <i class="space" on:click={handleSave}>
              <Fa icon={faCheck} color="green" />
            </i>
            <i class="space move" on:click={handleReset}>
              <Fa icon={faXmark} color="red" />
            </i>
          </span>
        {/if}
      </p>
    </div>
  </div>
</div>

<p class="content is-small">
  {description}
</p>

<style>
  .space {
    padding: 10px;
  }

  .move {
    margin-right: 40px;
  }
</style>
