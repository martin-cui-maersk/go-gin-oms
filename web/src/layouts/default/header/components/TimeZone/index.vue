<template>
  <Dropdown
    placement="bottom"
    :trigger="['hover']"
    :dropMenuList="timeZoneList"
    :selectedKeys="selectedKeys"
    @menu-event="handleMenuEvent"
    overlayClassName="app-locale-picker-overlay"
  >
    <span class="cursor-pointer flex items-center" style="margin-left: -10px">
      <Icon icon="ant-design:clock-circle-outlined" />
      <span v-if="showText" class="ml-1" style="margin-left: -5px; font-size: 13px">
        {{ getTimeZoneText }}
      </span>
    </span>
  </Dropdown>
</template>
<script lang="ts" setup>
  import type { TimeZoneType } from '#/config';
  import type { DropMenu } from '@/components/Dropdown';
  import { ref, watchEffect, unref, computed } from 'vue';
  import { Dropdown } from '@/components/Dropdown';
  import Icon from '@/components/Icon/Icon.vue';
  import { timeZoneList } from '@/settings/localeSetting';
  import { useLocaleStoreWithOut } from '@/store/modules/locale';

  defineOptions({ name: 'TimeZone' });

  const localeStore = useLocaleStoreWithOut();
  const getTimeZone = computed(() => localeStore.getTimeZone);

  const props = defineProps({
    /**
     * Whether to display text
     */
    showText: { type: Boolean, default: true },
    /**
     * Whether to refresh the interface when changing
     */
    reload: { type: Boolean },
  });

  const selectedKeys = ref<string[]>([]);

  const getTimeZoneText = computed(() => {
    const key = selectedKeys.value[0];
    if (!key) {
      return '';
    }
    return timeZoneList.find((item) => item.event === key)?.text;
  });

  watchEffect(() => {
    selectedKeys.value = [unref(getTimeZone)];
    console.log('watchEffect', selectedKeys.value);
  });

  async function toggleTimeZone(timeZone: TimeZoneType | string) {
    await localeStore.setLocaleInfo({ timeZone });
    selectedKeys.value = [timeZone as string];
    props.reload && location.reload();
  }

  function handleMenuEvent(menu: DropMenu) {
    if (unref(getTimeZone) === menu.event) {
      return;
    }
    toggleTimeZone(menu.event as string);
  }
</script>

<style lang="less">
  .app-locale-picker-overlay {
    .ant-dropdown-menu-item {
      min-width: 160px;
    }
  }
</style>
