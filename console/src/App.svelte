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

  const { t } = $i18n

  $menuLinks = [
    {
      path: '/key-manager',
      label: t('page.key-manager.menu-label'),
      component: KeyManager,
    },
    {
      path: '/history',
      label: t('page.history.menu-label'),
      component: History,
    },
    {
      path: '/settings',
      label: t('page.settings.menu-label'),
      component: Settings,
    },
  ]
  $menuActions = [
    {
      path: '/send-data',
      label: t('page.send-data.menu-label'),
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
    duration: 4000, // duration of progress bar tween to the `next` value
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
      component={NewUsers}
      meta={{ name: t('page.new-users.menu-label') }}
      primary={false}
    />
    <Route
      path="/register-key"
      component={RegisterKey}
      meta={{ name: t('page.register-key.menu-label') }}
      primary={false}
    />
    {#each $menuLinks as menuLink (menuLink.path)}
      <Route
        path={menuLink.path}
        component={menuLink.component}
        meta={{ name: menuLink.label }}
        primary={false}
      />
    {/each}
    {#each $menuActions as menuAction (menuAction.path)}
      <Route
        path={menuAction.path}
        component={menuAction.component}
        meta={{ name: menuAction.label }}
        primary={false}
      />
    {/each}
  </GlobalStyle>
</Router>

{#if $spinCount > 0}
  <Spinner />
{/if}

<SvelteToast options={toastOptions} />
