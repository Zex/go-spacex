package client

// TODO(zex)
// 1. configuable
// 2. better format

const (
  ApiScheme = "https"
  ApiHost = "api.spacexdata.com"
  ApiListRockets = "/v3/rockets"
  ApiGetRocket = "/v3/rockets/%s"
  ApiGetRoadster = "/v3/roadster"
  ApiListShips = "/v3/ships"
  ApiGetShip = "/v3/ships/%s"
  ApiListPayloads = "/v3/payloads"
  ApiGetPayload = "/v3/payload/%s"
  ApiListMissions = "/v3/missions"
  ApiGetMission = "/v3/mission/%s"
  ApiListDragons = "/v3/dragons"
  ApiGetDragon = "/v3/dragon/%s"
  ApiListLaunchPads = "/v3/launchpads"
  ApiGetLaunchPad = "/v3/launchpads/%s"
  ApiListLaunches = "/v3/launches"
  ApiGetLaunch = "/v3/launches/%s"
  ApiPastLaunches = "/v3/launches/past"
  ApiUpcomingLaunches = "/v3/launches/upcoming"
  ApiLatestLaunch = "/v3/launches/latest"
  ApiNextLaunch = "/v3/launches/next"
  ApiListLandPads = "/v3/landpads"
  ApiGetLandPad = "/v3/landpads/%s"
  ApiAboutUs = "/v3/info"
  ApiAboutApi = "/v3"
)
