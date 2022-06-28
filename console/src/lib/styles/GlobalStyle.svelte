<script>
  import { themeLight } from './theme-light'

  export let theme = 'light'

  let themeProps = {}

  $: {
    switch (theme) {
      case 'light':
        themeProps = themeLight
    }
  }
</script>

<slot />

<!--

--theme-primary
--theme-secondary
--theme-surface
--theme-background
--theme-on-primary
--theme-on-secondary
--theme-on-surface

-->
<style>
  :global(:root) {
    --box-sizing: border-box;

    background: #f4f6ff;
  }

  :global(body) {
    box-sizing: border-box;
    margin: 0;
    padding: 0;

    font-family: 'Work Sans';
  }

  /* :global(:root) {
    --color-primary: #14854f;
    --color-lightGrey: #d2d6dd;
    --color-grey: #747681;
    --color-darkGrey: #3f4144;
    --color-error: #d43939;
    --color-success: #28bd14;
    --grid-maxWidth: 120rem;
    --grid-gutter: 2rem;
    --font-size: 1.6rem;
    --font-family: 'Helvetica Neue', sans-serif; */
  /* } */
</style>
