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

<style>
  :global(:root) {
    --box-sizing: border-box;
    --font-family: 'Work Sans';

    background: #f4f6ff;
  }

  :global(body) {
    box-sizing: border-box;
    margin: 0;
    padding: 0;

    font-family: 'Work Sans';
  }

  :global(a),
  :global(a:hover),
  :global(a:focus),
  :global(a:visited),
  :global(a:active) {
    text-decoration: none;
    color: #0094ff;
  }

  :root {
    --toastContainerTop: auto;
    --toastContainerRight: auto;
    --toastContainerBottom: 30px;
    --toastContainerLeft: 30px;
  }
</style>
