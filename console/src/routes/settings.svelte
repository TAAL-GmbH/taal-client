<script>
  import { onMount } from "svelte";

  import Key from "./key.svelte";
  import Keys from "./keys.svelte";
  import Register from "./register.svelte";


  import Notifications from 'svelte-notifications';
  
  let keys;
  let key;

</script>

<main>
  <Keys {keys} />
  <Notifications>
    <Register/>
  </Notifications>
</main>
