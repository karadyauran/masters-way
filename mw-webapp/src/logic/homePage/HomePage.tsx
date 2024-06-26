import {NavigateFunction, useNavigate} from "react-router-dom";
import {homeAccessIds} from "cypress/accessIds/homeAccessIds";
import {observer} from "mobx-react-lite";
import {
  readManifestClicked,
  startForFreeFirstBlockClicked,
  startForFreeWhoWeAreBlockClicked,
  tryNowClicked,
} from "src/analytics/homeAnalytics";
import logoLight from "src/assets/mastersWayLogoLight.svg";
import {Button, ButtonType} from "src/component/button/Button";
import {Footer} from "src/component/footer/Footer";
import {LOGO_TEXT} from "src/component/header/Header";
import {HorizontalContainer} from "src/component/horizontalContainer/HorizontalContainer";
import {HorizontalGridContainer} from "src/component/horizontalGridContainer/HorizontalGridContainer";
import {Image} from "src/component/image/Image";
import {Link} from "src/component/link/Link";
import {HeadingLevel, Title} from "src/component/title/Title";
import {VerticalContainer} from "src/component/verticalContainer/VerticalContainer";
import {AuthDAL} from "src/dataAccessLogic/AuthDAL";
import {languageStore} from "src/globalStore/LanguageStore";
import {userStore} from "src/globalStore/UserStore";
import {AdvantageItem} from "src/logic/homePage/advantageItem/AdvantageItem";
import {SystemItem} from "src/logic/homePage/systemItem/SystemItem";
import {pages} from "src/router/pages";
import {LanguageService} from "src/service/LanguageService";
import styles from "src/logic/homePage/HomePage.module.scss";

/**
 * GetStarted button click handler
 */
const getStarted = (navigate: NavigateFunction, userUuid?: string) => {
  return userUuid
    ? navigate(pages.user.getPath({uuid: userUuid}))
    : AuthDAL.authGoogle();
};

/**
 * Home page
 */
export const HomePage = observer(() => {
  const {user} = userStore;
  const {language} = languageStore;
  const navigate = useNavigate();

  return (
    <>
      <div className={styles.container}>
        <HorizontalGridContainer className={styles.welcomeBlock}>
          <VerticalContainer className={styles.welcomeInfoBlock}>
            <Title
              level={HeadingLevel.h1}
              text={LanguageService.home.title[language]}
              className={styles.title}
              placeholder=""
              dataCy={homeAccessIds.welcomeBlock.title}
            />
            <p className={styles.titleDescription}>
              {LanguageService.home.description[language]}
            </p>
            <Button
              onClick={() => {
                startForFreeFirstBlockClicked();
                getStarted(navigate, user?.uuid);
              }}
              buttonType={ButtonType.SUPER_SPECIAL_BEAUTIFUL_BUTTON}
              value={LanguageService.home.startForFreeButton[language]}
              className={styles.getStartedButton}
              dataCy={homeAccessIds.welcomeBlock.startButton}
            />
          </VerticalContainer>

          <Image
            alt="video"
            src="https://drive.google.com/thumbnail?id=11wbqfxD6TDVneqyj7xidGIW4QPRSpMBt&sz=w1000"
            className={styles.welcomeImageBlock}
          />
        </HorizontalGridContainer>
      </div>

      <VerticalContainer className={styles.advantagesBlock}>
        <Title
          level={HeadingLevel.h2}
          text={LanguageService.home.advantages.advantagesTitle[language]}
          className={styles.advantagesTitle}
          placeholder=""
        />
        <p className={styles.advantagesDescription}>
          {LanguageService.home.advantages.advantagesDescription[language]}
        </p>
        <HorizontalContainer className={styles.advantages}>
          <AdvantageItem
            title={LanguageService.home.advantages.personalizedLearning.title[language]}
            description={LanguageService.home.advantages.personalizedLearning.description[language]}
          />
          <AdvantageItem
            title={LanguageService.home.advantages.comprehensiveControl.title[language]}
            description={LanguageService.home.advantages.comprehensiveControl.description[language]}
          />
          <AdvantageItem
            title={LanguageService.home.advantages.fastFeedback.title[language]}
            description={LanguageService.home.advantages.fastFeedback.description[language]}
          />
          <AdvantageItem
            title={LanguageService.home.advantages.visualizationOfProgress.title[language]}
            description={LanguageService.home.advantages.visualizationOfProgress.description[language]}
          />
          <AdvantageItem
            title={LanguageService.home.advantages.supportiveNetwork.title[language]}
            description={LanguageService.home.advantages.supportiveNetwork.description[language]}
          />
          <AdvantageItem
            title={LanguageService.home.advantages.timeSavingEfficiency.title[language]}
            description={LanguageService.home.advantages.timeSavingEfficiency.description[language]}
          />
        </HorizontalContainer>
      </VerticalContainer>

      <VerticalContainer className={styles.buildSystemBlock}>
        <Title
          level={HeadingLevel.h2}
          text={LanguageService.home.systemBlock.title[language]}
          className={styles.systemBlockTitle}
          placeholder=""
        />
        <p className={styles.systemBlockDescription}>
          {LanguageService.home.systemBlock.description[language]}
        </p>
        <HorizontalContainer className={styles.systemItems}>
          <SystemItem
            title={LanguageService.home.systemBlock.createWay.title[language]}
            description={LanguageService.home.systemBlock.createWay.description[language]}
          />
          <SystemItem
            title={LanguageService.home.systemBlock.setGoalAndMetrics.title[language]}
            description={LanguageService.home.systemBlock.setGoalAndMetrics.description[language]}
          />
          <SystemItem
            title={LanguageService.home.systemBlock.trackProgress.title[language]}
            description={LanguageService.home.systemBlock.trackProgress.description[language]}
          />
        </HorizontalContainer>

        <Button
          onClick={() => {
            startForFreeWhoWeAreBlockClicked();
            getStarted(navigate, user?.uuid);
          }}
          buttonType={ButtonType.SUPER_SPECIAL_BEAUTIFUL_BUTTON}
          value={LanguageService.home.getStartedButton[language]}
          className={styles.getStartedButton}
        />
      </VerticalContainer>

      <VerticalContainer className={styles.aboutProjectBlock}>
        <Title
          level={HeadingLevel.h2}
          text={LanguageService.home.aboutAppBlock.title[language]}
          className={styles.aboutAppBlockTitle}
          placeholder=""
        />
        <HorizontalContainer className={styles.aboutProjectItems}>
          <Image
            alt={LOGO_TEXT}
            src={logoLight}
            className={styles.aboutProjectImage}
          />
          <VerticalContainer className={styles.aboutProjectDescription}>

            <p className={styles.aboutAppBlockDescription}>
              {LanguageService.home.aboutAppBlock.description[language]}
            </p>

            <HorizontalContainer className={styles.aboutAppBlockButtons}>
              <Button
                onClick={() => {
                  tryNowClicked();
                  getStarted(navigate, user?.uuid);
                }}
                buttonType={ButtonType.SUPER_SPECIAL_BEAUTIFUL_BUTTON}
                value={LanguageService.home.getStartedButton[language]}
                className={styles.aboutAppGetStartedButton}
              />
              <Link
                path={LanguageService.home.aboutAppBlock.manifestLink[language]}
                className={styles.manifestoLink}
                onClick={readManifestClicked}
                isNewTab
              >
                {LanguageService.home.aboutAppBlock.readManifesto[language]}
              </Link>
            </HorizontalContainer>
          </VerticalContainer>
        </HorizontalContainer>

      </VerticalContainer>

      <Footer language={language} />
    </>
  );

});
