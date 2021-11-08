package fetchers

import (
	. "github.com/onsi/gomega"
	"testing"

	"github.com/derbexuk/wurzel/harvester/feeds/config"
)

func TestSimplePaths(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : XML
frequency : 60m
feed source : http://api.erg.kcl.ac.uk/AirQuality/Hourly/MonitoringIndex/
sub feeds :
  SiteCode :
    - CR5
paths :
  -
    source : ./HourlyAirQualityIndex/LocalAuthority/Site/species/attrs/AirQualityBand
    destination : /environment/london/|SiteCode|/air-quality/band
  -
    source : ./HourlyAirQualityIndex/LocalAuthority/Site/species/attrs/AirQualityIndex
    destination : /environment/london/|SiteCode|/air-quality/index
`
	data := `
<?xml version="1.0" encoding="utf-8"?><HourlyAirQualityIndex TimeToLive="55"><LocalAuthority LocalAuthorityName="Croydon" LocalAuthorityCode="8" LaCentreLatitude="51.372361" LaCentreLongitude="-0.100401" LaCentreLongitudeWGS84="-11176.588195" LaCentreLatitudeWGS84="6687426.268766"><Site LatitudeWGS84="6694381.70052" LongitudeWGS84="-13704.5425116" Longitude="-0.12311" Latitude="51.411349" SiteType="Kerbside" SiteCode="CR5" SiteName="Croydon - Norbury" BulletinDate="2017-03-14 10:00:00"><species SpeciesName="Nitrogen Dioxide" SpeciesCode="NO2" AirQualityIndex="1" AirQualityBand="Low" IndexSource="Measurement"/></Site></LocalAuthority></HourlyAirQualityIndex>
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &XMLFetcher{Config: fc}
	res := feed.process([]byte(data), map[string]string{"SiteCode": "CR5"})
	Expect(res["/environment/london/CR5/air-quality/band"]).To(Equal("Low"))

	data = `
<?xml version="1.0" encoding="utf-8"?><HourlyAirQualityIndex TimeToLive="55"><LocalAuthority LocalAuthorityName="Croydon" LocalAuthorityCode="8" LaCentreLatitude="51.372361" LaCentreLongitude="-0.100401" LaCentreLongitudeWGS84="-11176.588195" LaCentreLatitudeWGS84="6687426.268766"><Site LatitudeWGS84="6694381.70052" LongitudeWGS84="-13704.5425116" Longitude="-0.12311" Latitude="51.411349" SiteType="Kerbside" SiteCode="CR5" SiteName="Croydon - Norbury" BulletinDate="2017-03-14 10:00:00"></Site></LocalAuthority></HourlyAirQualityIndex>
`
	res = feed.process([]byte(data), map[string]string{"SiteCode": "CR5"})
	Expect(res["/environment/london/CR5/air-quality/band"]).To(Equal(""))
}

func TestEvents(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : XML
frequency : 60s
feed source : http://football.api.press.net/v1.5/match/events/PqX9xW7wXDB2/3979877
paths :
  -
    source : matchEvents/isResult
    destination : /sports/epl/match/12345/result
  -
    source : matchEvents/events
    destination : /sports/epl/match/12345/events
`
	data := `
<?xml version="1.0" encoding="utf-8"?>
<matchEvents>
    <matchMinutes>90</matchMinutes>
    <isResult>Yes</isResult>
    <teams>
        <homeTeam teamID="9">Liverpool</homeTeam>
        <awayTeam teamID="8">Everton</awayTeam>
    </teams>
    <events>
        <event teamID="" eventID="">
            <eventType>timeline</eventType>
            <matchTime>0:00</matchTime>
            <eventTime>0</eventTime>
            <normalTime>0:00</normalTime>
            <addedTime>0:00</addedTime>
            <players>
                <player1 playerID="" teamID="" teamName="" />
                <player2 playerID="" teamID="" teamName="" />
            </players>
            <reason />
            <how />
            <whereFrom />
            <whereTo />
            <type />
            <distance />
            <outcome />
        </event>
        <event teamID="" eventID="">
            <eventType>timeline</eventType>
            <matchTime>(90 +0:42)</matchTime>
            <eventTime>90</eventTime>
            <normalTime>90:00</normalTime>
            <addedTime>0:42</addedTime>
            <players>
                <player1 playerID="" teamID="" teamName="" />
                <player2 playerID="" teamID="" teamName="" />
            </players>
            <reason />
            <how />
            <whereFrom />
            <whereTo />
            <type />
            <distance />
            <outcome />
        </event>

        <event teamID="9" eventID="23458032">
            <eventType>goal kick</eventType>
            <matchTime>(17)</matchTime>
            <eventTime>17</eventTime>
            <normalTime>16:28</normalTime>
            <addedTime>0:00</addedTime>
            <players>
                <player1 playerID="322905" teamID="9" teamName="Liverpool">Simon Mignolet</player1>
                <player2 playerID="" teamID="" teamName="" />
            </players>
            <reason />
            <how />
            <whereFrom />
            <whereTo />
            <type>Long</type>
            <distance />
            <outcome />
        </event>
        <event teamID="9" eventID="23458033">
            <eventType>throw in</eventType>
            <matchTime>(18)</matchTime>
            <eventTime>18</eventTime>
            <normalTime>17:28</normalTime>
            <addedTime>0:00</addedTime>
            <players>
                <player1 playerID="247364" teamID="9" teamName="Liverpool">James Milner</player1>
                <player2 playerID="" teamID="" teamName="" />
            </players>
            <reason />
            <how />
            <whereFrom />
            <whereTo />
            <type>Attacking</type>
            <distance />
            <outcome />
        </event>
        <event teamID="9" eventID="23458034">
            <eventType>throw in</eventType>
            <matchTime>(18)</matchTime>
            <eventTime>18</eventTime>
            <normalTime>17:41</normalTime>
            <addedTime>0:00</addedTime>
            <players>
                <player1 playerID="247364" teamID="9" teamName="Liverpool">James Milner</player1>
                <player2 playerID="" teamID="" teamName="" />
            </players>
            <reason />
            <how />
            <whereFrom />
            <whereTo />
            <type>Attacking</type>
            <distance />
            <outcome />
        </event>
        <event teamID="9" eventID="23458035">
            <eventType>shot off target</eventType>
            <matchTime>(19)</matchTime>
            <eventTime>19</eventTime>
            <normalTime>18:42</normalTime>
            <addedTime>0:00</addedTime>
            <players>
                <player1 playerID="443680" teamID="9" teamName="Liverpool">Phillippe Coutinho</player1>
                <player2 playerID="" teamID="" teamName="" />
            </players>
            <reason />
            <how>Right Foot</how>
            <whereFrom>Centre Penalty Area</whereFrom>
            <whereTo />
            <type />
            <distance />
            <outcome>Blocked</outcome>
        </event>
    </events>
</matchEvents>
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &XMLFetcher{Config: fc}
	res := feed.process([]byte(data), nil)

	Expect(res["/sports/epl/match/12345/events"]).Should(ContainSubstring(`"-playerID": "443680"`))
	Expect(res["/sports/epl/match/12345/result"]).To(Equal("Yes"))

	upcoming_match := `<?xml version="1.0" encoding="utf-8"?>
<matchEvents>
    <isResult>No</isResult>
    <teams>
        <homeTeam teamID="41">Watford</homeTeam>
        <awayTeam teamID="11">Man City</awayTeam>
    </teams>
    <events />
</matchEvents>`

	feed = &XMLFetcher{Config: fc}
	res = feed.process([]byte(upcoming_match), nil)

	Expect(res["/sports/epl/match/12345/events"]).To(Equal(""))
	Expect(res["/sports/epl/match/12345/result"]).To(Equal("No"))

}

func TestResetEvents(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : XML
frequency : 60s
feed source : http://football.api.press.net/v1.5/match/events/PqX9xW7wXDB2/3979877
paths :
  -
    source : matchEvents/isResult
    destination : /sports/epl/match/12345/result
  -
    source : matchEvents/events
    destination : /sports/epl/match/12345/events
`
	reset_event := `<?xml version="1.0" encoding="utf-8"?><matchEvents><isResult>No</isResult><events>RESET</events></matchEvents>`

	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))

	feed := &XMLFetcher{Config: fc}
	res := feed.process([]byte(reset_event), nil)

	Expect(res["/sports/epl/match/12345/events"]).To(Equal("RESET"))
	Expect(res["/sports/epl/match/12345/result"]).To(Equal("No"))
}

func TestCompositePaths(t *testing.T) {
	RegisterTestingT(t)

	yaml := `
---
format : XML
frequency : 60m
feed source : http://api.erg.kcl.ac.uk/AirQuality/Hourly/MonitoringIndex/
paths :
  -
    source : fixtures
    destination : /sports/pml/fixtures
`
	data := `
<?xml version="1.0" encoding="utf-8"?>
<fixtures>
    <fixture matchID="3925337" date="21/05/2017" koTime="15:00" originalID="">
        <competition competitionID="100" seasonID="4255">Premier League 16/17</competition>
        <stage stageNumber="1" stageType="League">
        </stage>
        <round roundNumber="1">League</round>
        <leg>1</leg>
        <homeTeam teamID="65">Swansea</homeTeam>
        <awayTeam teamID="42">West Brom</awayTeam>
        <venue venueID="79">Liberty Stadium</venue>
        <referee>
        </referee>
    </fixture>
    <fixture matchID="3925338" date="21/05/2017" koTime="15:00" originalID="">
        <competition competitionID="100" seasonID="4255">Premier League 16/17</competition>
        <stage stageNumber="1" stageType="League">
        </stage>
        <round roundNumber="1">League</round>
        <leg>1</leg>
        <homeTeam teamID="41">Watford</homeTeam>
        <awayTeam teamID="11">Man City</awayTeam>
        <venue venueID="80">Vicarage Road Stadium</venue>
        <referee>
        </referee>
    </fixture>
</fixtures>
`
	fc := config.FeedConfig{}
	fc.Populate([]byte(yaml))
	feed := &XMLFetcher{Config: fc}
	res := feed.process([]byte(data), nil)

	Expect(res["/sports/pml/fixtures"]).Should(ContainSubstring(`"venue": {"#content": "Vicarage Road Stadium", "-venueID": "80"}`))
}
