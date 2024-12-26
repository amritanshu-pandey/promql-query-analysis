package dashboardQueries

type Query struct {
	queryString    string
	datasourceType string
	panelId        string
}

type Dashboard struct {
	Name    string
	UID     string
	Queries []Query
}

func NewDashboardQueries(gc *GrafanaClient, dashboardUID string) (Dashboard, error) {
	dashboard, err := gc.Dashboard(dashboardUID)
	if err != nil {
		return Dashboard{}, err
	}
	return dashboard, nil
}
