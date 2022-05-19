<script>
  import Keys from "./keys.svelte";
  import Register from "./register.svelte";
  import Notifications from 'svelte-notifications';
  
  let keys;

</script>

<main>
  <Keys {keys} />
  <Notifications>
    <Register/>
  </Notifications>
</main>
