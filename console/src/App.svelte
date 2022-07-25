<script>
  import { Router, Route } from 'svelte-navigator'
  import { SvelteToast } from '@zerodevx/svelte-toast'

  import { menuLinks, menuActions, mediaSize } from './lib/stores'
  import { query } from './lib/actions'
  import { medium, large } from './lib/styles/breakpoints'
  import GlobalStyle from './lib/styles/GlobalStyle.svelte'
  import NewUsers from './routes/new-users/index.svelte'
  import RegisterKey from './routes/register-key/index.svelte'
  import KeyManager from './routes/key-manager/index.svelte'
  import History from './routes/history/index.svelte'
  import Settings from './routes/settings/index.svelte'
  import SendData from './routes/send-data/index.svelte'
  import Spinner from './lib/components/spinner/index.svelte'
  import { spinCount } from './lib/stores'
  import i18n from './lib/i18n'

  $menuLinks = [
    {
      path: '/key-manager',
      label: $i18n.t('page.key-manager.menu-label'),
      component: KeyManager,
    },
    {
      path: '/history',
      label: $i18n.t('page.history.menu-label'),
      component: History,
    },
    {
      path: '/settings',
      label: $i18n.t('page.settings.menu-label'),
      component: Settings,
    },
  ]
  $menuActions = [
    {
      path: '/send-data',
      label: $i18n.t('page.send-data.menu-label'),
      component: SendData,
    },
  ]

  $: isMedium = query(medium)
  $: isLarge = query(large)
  $: {
    $mediaSize = $isLarge ? 'large' : $isMedium ? 'medium' : 'small'
  }

  export let url = ''

  const toastOptions = {
    duration: 3000, // duration of progress bar tween to the `next` value
    pausable: true, // pause progress bar tween on mouse hover
    dismissable: true, // allow dismiss with close button
    reversed: false, // insert new toast to bottom of stack
    intro: { y: 192 },
  }
</script>

<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com" />
  <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="true" />
  <link
    href="https://fonts.googleapis.com/css2?family=Work+Sans:wght@400;500;600;700&display=swap"
    rel="stylesheet"
  />
</svelte:head>

<Router {url}>
  <GlobalStyle>
    <Route
      path="/"
      meta={{ name: $i18n.t('page.new-users.menu-label') }}
      primary={false}
    >
      <NewUsers />
    </Route>
    <Route
      path="/register-key"
      meta={{ name: $i18n.t('page.register-key.menu-label') }}
      primary={false}
    >
      <RegisterKey />
    </Route>
    {#each $menuLinks as menuLink (menuLink.path)}
      <Route
        path={menuLink.path}
        meta={{ name: menuLink.label }}
        primary={false}
      >
        <svelte:component this={menuLink.component} />
      </Route>
    {/each}
    {#each $menuActions as menuAction (menuAction.path)}
      <Route
        path={menuAction.path}
        meta={{ name: menuAction.label }}
        primary={false}
      >
        <svelte:component this={menuAction.component} />
      </Route>
    {/each}
  </GlobalStyle>
</Router>

{#if $spinCount > 0}
  <Spinner />
{/if}

<SvelteToast options={toastOptions} />
