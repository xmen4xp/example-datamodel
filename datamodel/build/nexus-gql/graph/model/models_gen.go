// Code generated by github.com/vmware-tanzu/graph-framework-for-microservices/gqlgen, DO NOT EDIT.

package model

type NexusGraphqlResponse struct {
	Code         *int    `json:"Code"`
	Message      *string `json:"Message"`
	Data         *string `json:"Data"`
	Last         *string `json:"Last"`
	TotalRecords *int    `json:"TotalRecords"`
}

type TimeSeriesData struct {
	Code         *int    `json:"Code"`
	Message      *string `json:"Message"`
	Data         *string `json:"Data"`
	Last         *string `json:"Last"`
	TotalRecords *int    `json:"TotalRecords"`
}

type ConfigConfig struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	User         []*UserUser            `json:"User"`
	Event        []*EventEvent          `json:"Event"`
}

type EvaluationEvaluation struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	Quiz         []*QuizQuiz            `json:"Quiz"`
}

type EventEvent struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	Description  *string                `json:"Description"`
	MeetingLink  *string                `json:"MeetingLink"`
	Time         *string                `json:"Time"`
	Public       *bool                  `json:"Public"`
}

type InterestInterest struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	Name         *string                `json:"Name"`
}

type QuizQuiz struct {
	Id                      *string                     `json:"Id"`
	ParentLabels            map[string]interface{}      `json:"ParentLabels"`
	Question                []*QuizquestionQuizQuestion `json:"Question"`
	DefaultScorePerQuestion *int                        `json:"DefaultScorePerQuestion"`
}

type QuizchoiceQuizChoice struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	Choice       *string                `json:"Choice"`
	Hint         *string                `json:"Hint"`
	PictureName  *string                `json:"PictureName"`
	Answer       *bool                  `json:"Answer"`
}

type QuizquestionQuizQuestion struct {
	Id                *string                 `json:"Id"`
	ParentLabels      map[string]interface{}  `json:"ParentLabels"`
	Choice            []*QuizchoiceQuizChoice `json:"Choice"`
	Question          *string                 `json:"Question"`
	Hint              *string                 `json:"Hint"`
	Format            *string                 `json:"Format"`
	Score             *int                    `json:"Score"`
	AnimationFilePath *string                 `json:"AnimationFilePath"`
	PictureFilePath   *string                 `json:"PictureFilePath"`
}

type RootRoot struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	Tenant       []*TenantTenant        `json:"Tenant"`
	Evaluation   *EvaluationEvaluation  `json:"Evaluation"`
}

type RuntimeRuntime struct {
	Id           *string                   `json:"Id"`
	ParentLabels map[string]interface{}    `json:"ParentLabels"`
	User         []*RuntimeuserRuntimeUser `json:"User"`
}

type RuntimeanswerRuntimeAnswer struct {
	Id             *string                `json:"Id"`
	ParentLabels   map[string]interface{} `json:"ParentLabels"`
	Answer         *QuizchoiceQuizChoice  `json:"Answer"`
	ProvidedAnswer *string                `json:"ProvidedAnswer"`
}

type RuntimeevaluationRuntimeEvaluation struct {
	Id           *string                   `json:"Id"`
	ParentLabels map[string]interface{}    `json:"ParentLabels"`
	Quiz         []*RuntimequizRuntimeQuiz `json:"Quiz"`
}

type RuntimequizRuntimeQuiz struct {
	Id           *string                       `json:"Id"`
	ParentLabels map[string]interface{}        `json:"ParentLabels"`
	Quiz         *QuizQuiz                     `json:"Quiz"`
	Answers      []*RuntimeanswerRuntimeAnswer `json:"Answers"`
}

type RuntimeuserRuntimeUser struct {
	Id           *string                             `json:"Id"`
	ParentLabels map[string]interface{}              `json:"ParentLabels"`
	User         *UserUser                           `json:"User"`
	Evaluation   *RuntimeevaluationRuntimeEvaluation `json:"Evaluation"`
}

type TenantTenant struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	Interest     []*InterestInterest    `json:"Interest"`
	Config       *ConfigConfig          `json:"Config"`
	Runtime      *RuntimeRuntime        `json:"Runtime"`
}

type UserUser struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	Wanna        []*WannaWanna          `json:"Wanna"`
	Username     *string                `json:"Username"`
	Mail         *string                `json:"Mail"`
	FirstName    *string                `json:"FirstName"`
	LastName     *string                `json:"LastName"`
	Password     *string                `json:"Password"`
	Realm        *string                `json:"Realm"`
}

type WannaWanna struct {
	Id           *string                `json:"Id"`
	ParentLabels map[string]interface{} `json:"ParentLabels"`
	Interest     *InterestInterest      `json:"Interest"`
	Name         *string                `json:"Name"`
}
