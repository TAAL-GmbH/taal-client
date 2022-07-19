<script>
  import { SvelteToast } from '@zerodevx/svelte-toast'

  import AppOld from './AppOld.svelte'
  import AppNew from './AppNew.svelte'

  let old = false

  const toastOptions = {
    duration: 4000, // duration of progress bar tween to the `next` value
    pausable: true, // pause progress bar tween on mouse hover
    dismissable: true, // allow dismiss with close button
    reversed: false, // insert new toast to bottom of stack
    intro: { y: 192 },
  }
</script>

{#if old}
  <AppOld />
{:else}
  <AppNew />
  <SvelteToast options={toastOptions} />
{/if}
