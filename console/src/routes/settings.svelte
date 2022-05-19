<script>
  import { onMount } from "svelte";

  import Key from "./key.svelte";
  import Keys from "./keys.svelte";
  import Register from "./register.svelte";
  
  let keys;
  let key;

</script>

<main>
  <Keys {keys} />
  <Register/>
</main>
