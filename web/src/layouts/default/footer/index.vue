<template>
  <Layout.Footer :class="prefixCls" v-if="getShowLayoutFooter" ref="footerRef">
    <div :class="`${prefixCls}__links`">
      <a @click="openWindow(OFFICIAL_SITE_URL)">{{ t('layout.footer.siteName') }}</a>
    </div>
    <div>Copyright &copy; 2024 Maersk E-Partner</div>
  </Layout.Footer>
</template>
<script lang="ts" setup>
  import { computed, unref, ref } from 'vue';
  import { Layout } from 'ant-design-vue';

  // import { GithubFilled } from '@ant-design/icons-vue';

  import { OFFICIAL_SITE_URL } from '@/settings/siteSetting';
  import { openWindow } from '@/utils';

  import { useI18n } from '@/hooks/web/useI18n';
  import { useRootSetting } from '@/hooks/setting/useRootSetting';
  import { useRouter } from 'vue-router';
  import { useDesign } from '@/hooks/web/useDesign';
  import { useLayoutHeight } from '../content/useContentViewHeight';

  defineOptions({ name: 'LayoutFooter' });

  const { t } = useI18n();
  const { getShowFooter } = useRootSetting();
  const { currentRoute } = useRouter();
  const { prefixCls } = useDesign('layout-footer');

  const footerRef = ref<ComponentRef>(null);
  const { setFooterHeight } = useLayoutHeight();

  const getShowLayoutFooter = computed(() => {
    if (unref(getShowFooter)) {
      const footerEl = unref(footerRef)?.$el;
      setFooterHeight(footerEl?.offsetHeight || 0);
    } else {
      setFooterHeight(0);
    }
    return unref(getShowFooter) && !unref(currentRoute).meta?.hiddenFooter;
  });
</script>
<style lang="less" scoped>
  @prefix-cls: ~'@{namespace}-layout-footer';

  @normal-color: rgba(0, 0, 0, 0.45);

  @hover-color: rgba(0, 0, 0, 0.85);

  .@{prefix-cls} {
    // 页脚固定高度
    height: 75px;
    color: @normal-color;
    text-align: center;

    &__links {
      margin-bottom: 8px;

      a {
        color: @normal-color;

        &:hover {
          color: @hover-color;
        }
      }
    }

    &__github {
      margin: 0 30px;

      &:hover {
        color: @hover-color;
      }
    }
  }
</style>
