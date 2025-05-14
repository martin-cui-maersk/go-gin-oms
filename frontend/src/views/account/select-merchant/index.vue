<template>
  <PageWrapper class="list-card" title="商户列表">
    <template #headerContent>点击商户进行切换</template>
    <div class="list-card__content">
      <List :loading="cardLoading">
        <Row :gutter="16">
          <template v-for="item in cardList" :key="item.merchantId">
            <Col :xl="6" :md="8" :sm="8">
              <List.Item>
                <Card
                  :style="
                    userInfo.merchantId === item.merchantId
                      ? 'border: 1px solid rgb(24, 144, 255)'
                      : ''
                  "
                  :hoverable="true"
                  class="list-card__card"
                  @click="clickMerchant(item)"
                  :border="true"
                >
                  <div class="list-card__card-title">
                    <Icon class="icon" icon="ant-design:shop-twotone" color="#1890ff" />
                    {{ item.merchantName }}
                  </div>
                  <div :class="`list-card__card-detail`">
                    商户编码: {{ item.merchantCode }} <br />
                  </div>
                </Card>
              </List.Item>
            </Col>
          </template>
        </Row>
      </List>
    </div>
  </PageWrapper>
</template>
<script lang="ts" setup>
  import Icon from '@/components/Icon/Icon.vue';
  import { PageWrapper } from '@/components/Page';
  import { Card, Row, Col, List } from 'ant-design-vue';
  import { getUserMerchantList, userChangeMerchant } from '@/api/basic/user';
  import { ref, onMounted, h } from 'vue';
  import { useUserStore } from '@/store/modules/user';
  import { useMessage } from '@/hooks/web/useMessage';
  import { useI18n } from '@/hooks/web/useI18n';
  // import { useGo } from '@/hooks/web/usePage';
  import { PageEnum } from '@/enums/pageEnum';
  // import { useRouter } from 'vue-router';
  import { router } from '@/router';

  const userStore = useUserStore();
  const userInfo = userStore.getUserInfo;
  // const go = useGo();

  function clickMerchant(item) {
    console.log(item);
    if (userInfo.merchantId === item.merchantId) {
      return false;
    }
    const { createConfirm } = useMessage();
    const { t } = useI18n();
    createConfirm({
      iconType: 'info',
      title: () => h('span', t('account.merchant.tip')),
      content: () => h('span', t('account.merchant.confirmMessage')),
      onOk: async () => {
        await changeMerchant(item.merchantId);
      },
    });
  }

  async function changeMerchant(id) {
    return new Promise((resolve, reject) => {
      userChangeMerchant(id).then(async (res) => {
        if (res) {
          // go(PageEnum.BASE_HOME);
          await router.replace(PageEnum.BASE_HOME).then(() => {
            window.location.reload();
          });
          resolve(true);
        } else {
          reject(false);
        }
      });
    });
  }

  let cardLoading = ref(true);
  let cardList = ref([]);
  onMounted(() => {
    getUserMerchantList().then((res) => {
      cardList.value = res;
      cardLoading.value = false;
    });
  });
</script>
<style lang="less" scoped>
  .list-card {
    &__link {
      margin-top: 10px;
      font-size: 14px;

      a {
        margin-right: 30px;
      }

      span {
        margin-left: 5px;
      }
    }

    &__card {
      width: 100%;
      margin-bottom: -8px;

      .ant-card-body {
        padding: 16px;
      }

      &-title {
        margin-bottom: 5px;
        color: @text-color-base;
        font-size: 16px;
        font-weight: 500;

        .icon {
          margin-top: -5px;
          margin-right: 10px;
          font-size: 38px !important;
        }
      }

      &-detail {
        padding-top: 10px;
        padding-left: 30px;
        color: @text-color-secondary;
        font-size: 14px;
      }
    }
  }
</style>
