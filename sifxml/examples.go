package sifxml

var test_ActivityExample1 = `		<Activity RefId="C27E1FCF-C163-485F-BEF0-F36F18A0493A">
			<Title>Shakespeare Essay - Much Ado About Nothing</Title>
			<Preamble>This is a very funny comedy - students should have passing familiarity with Shakespeare</Preamble>
			<LearningStandards>
				<LearningStandardItemRefId>9DB15CEA-B2C5-4F66-94C3-7D0A0CAEDDA4</LearningStandardItemRefId>
			</LearningStandards>
			<SourceObjects>
				<SourceObject SIF_RefObject="Lesson">A71ADBD3-D93D-A64B-7166-E420D50EDABC</SourceObject>
			</SourceObjects>
			<Points>50</Points>
			<ActivityTime>
				<CreationDate>2002-06-15</CreationDate>
				<Duration Units="minute">30</Duration>
				<StartDate>2002-09-10</StartDate>
				<FinishDate>2002-09-12</FinishDate>
				<DueDate>2002-09-12</DueDate>
			</ActivityTime>
			<AssessmentRefId>03EDB29E-8116-B450-0435-FA87E42A0AD2</AssessmentRefId>
			<MaxAttemptsAllowed>3</MaxAttemptsAllowed>
			<ActivityWeight>5</ActivityWeight>
			<Evaluation EvaluationType="Inline">
				<Description>Students should be able to correctly identify all major characters.</Description>
			</Evaluation>
			<LearningResources>
				<LearningResourceRefId>B7337698-BF6D-B193-7F79-A07B87211B93</LearningResourceRefId>
			</LearningResources>
		</Activity>
`
var test_ActivityExample2 = `		<Activity RefId="C27E1FCF-C163-485F-BEF0-F36F18A0493A">
			<Title>Shakespeare Essay - Much Ado About Nothing</Title>
			<Preamble>This is a very funny comedy - students should have passing familiarity with Shakespeare</Preamble>
			<SoftwareRequirementList>
				<SoftwareRequirement>
					<SoftwareTitle>Flash Player</SoftwareTitle>
					<Version>9.0</Version>
					<Vendor>Adobe</Vendor>
				</SoftwareRequirement>
				<SoftwareRequirement>
					<SoftwareTitle>Python</SoftwareTitle>
					<Version>3.0</Version>
					<OS>Linux</OS>
				</SoftwareRequirement>
			</SoftwareRequirementList>
			<LearningStandards>
				<LearningStandardItemRefId>9DB15CEA-B2C5-4F66-94C3-7D0A0CAEDDA4</LearningStandardItemRefId>
			</LearningStandards>
			<SourceObjects>
				<SourceObject SIF_RefObject="Lesson">A71ADBD3-D93D-A64B-7166-E420D50EDABC</SourceObject>
			</SourceObjects>
			<Points>50</Points>
			<ActivityTime>
				<CreationDate>2002-06-15</CreationDate>
				<Duration Units="minute">30</Duration>
				<StartDate>2002-09-10</StartDate>
				<FinishDate>2002-09-12</FinishDate>
				<DueDate>2002-09-12</DueDate>
			</ActivityTime>
			<AssessmentRefId>03EDB29E-8116-B450-0435-FA87E42A0AD2</AssessmentRefId>
			<MaxAttemptsAllowed>3</MaxAttemptsAllowed>
			<ActivityWeight>5</ActivityWeight>
			<Evaluation EvaluationType="Inline">
				<Description>Students should be able to correctly identify all major characters.</Description>
			</Evaluation>
			<LearningResources>
				<LearningResourceRefId>B7337698-BF6D-B193-7F79-A07B87211B93</LearningResourceRefId>
			</LearningResources>
		</Activity>
`
var test_AggregateCharacteristicInfo = `    <AggregateCharacteristicInfo RefId="06AF1D69-FB06-4AB3-8898-3505714F18FB">
      <Description>Sex</Description>
      <Definition>Gender or Sex</Definition>
      <ElementName>StudentPersonal/PersonInfo/Demographics/Sex</ElementName>
    </AggregateCharacteristicInfo>
`
var test_AggregateStatisticFact = `    <AggregateStatisticFact RefId="3A4E2822-C696-433B-812B-49DEAE557E41">
      <AggregateStatisticInfoRefId>91A4209F-7F4F-4F4B-9DC8-D21CEFD1DC2F</AggregateStatisticInfoRefId>
      <Characteristics>
        <AggregateCharacteristicInfoRefId>0B972D86-44EE-4D79-94F5-FF1C3131772D</AggregateCharacteristicInfoRefId>
      </Characteristics>
      <Excluded>No</Excluded>
      <Value>27</Value>
    </AggregateStatisticFact>
`
var test_AggregateStatisticInfo = `    <AggregateStatisticInfo RefId="A5DECBD3-161B-4F5D-9F46-EBA64C87B002">
      <StatisticName>Read Proficiency L1</StatisticName>
      <CalculationRule Type="Description">Number of students scoring at Reading Proficiency Level 1</CalculationRule>
      <ApprovalDate>2002-09-01</ApprovalDate>
      <ExpirationDate>2006-06-30</ExpirationDate>
      <ExclusionRules>
        <ExclusionRule Type="SampleSize">N less than 10</ExclusionRule>
      </ExclusionRules>
      <Source>SIS</Source>
      <Location Type="School">
        <LocationName>Green River High School</LocationName>
        <LocationRefId SIF_RefObject="SchoolInfo">279BCEE5-515E-4C1A-AC3A-765D1F069BC3</LocationRefId>
      </Location>
      <Measure>Count</Measure>
    </AggregateStatisticInfo>
`
var test_CalendarDate = `    <CalendarDate CalendarDateRefId="B5739375-800A-C4CC-6385-0BB2754114AA">
      <Date>2007-08-31</Date>
      <CalendarSummaryRefId>A8A34E56-1C97-345C-0A4E-11BB113C2753</CalendarSummaryRefId>
      <SchoolInfoRefId>B7A34E56-1C97-345C-0A4E-11BB112B2753</SchoolInfoRefId>
      <SchoolYear>2007</SchoolYear>
      <CalendarDateType>
        <Code>INST</Code>
      </CalendarDateType>
      <StudentAttendance>
        <CountsTowardAttendance>Yes</CountsTowardAttendance>
        <AttendanceValue>1.0</AttendanceValue>
      </StudentAttendance>
      <TeacherAttendance>
        <CountsTowardAttendance>Yes</CountsTowardAttendance>
        <AttendanceValue>1.0</AttendanceValue>
      </TeacherAttendance>
    </CalendarDate>

`
var test_CalendarSummary = `    <CalendarSummary RefId="B5739375-800A-C4CC-6385-0BB2754114AA">
      <SchoolInfoRefId>B7A34E56-1C97-345C-0A4E-11BB112B2753</SchoolInfoRefId>
      <SchoolYear>2005</SchoolYear>
          <LocalId>123321A</LocalId>
          <Description>Elementary School Calendar</Description>
          <DaysInSession>180</DaysInSession>
          <StartDate>2009-01-01</StartDate>
          <EndDate>2009-12-31</EndDate>
          <FirstInstructionDate>2009-01-02</FirstInstructionDate>
          <LastInstructionDate>2009-12-20</LastInstructionDate>
          <GraduationDate>2009-12-10</GraduationDate>
          <InstructionalMinutes>64800</InstructionalMinutes>
          <MinutesPerDay>360</MinutesPerDay>
          <YearLevels>
            <YearLevel>
              <Code>10</Code>
            </YearLevel>
            <YearLevel>
              <Code>11</Code>
            </YearLevel>
            <YearLevel>
              <Code>12</Code>
            </YearLevel>
          </YearLevels>
        </CalendarSummary>

`
var test_EquipmentInfoExample = `    <EquipmentInfo RefId="B5739375-800A-C4CC-6385-0BB2754114AA">
      <Name>Overhead Projector 3</Name>
      <LocalId>OHP3</LocalId>
      <EquipmentType>OverheadProjector</EquipmentType>
      <SIF_RefId SIF_RefObject="SchoolInfo">B7A34E56-1C97-345C-0A4E-11BB112B2753</SIF_RefId>
    </EquipmentInfo>
    

`
var test_GradingAssignment = `    <GradingAssignment RefId="359D7510-1AD0-A9D7-A8C3-DAD0A85103A2">	
      <TeachingGroupRefId>D0A0A27A-D0A8-510A-D9D7-5101A8C3DA39</TeachingGroupRefId>
      <SchoolInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</SchoolInfoRefId>
      <GradingCategory>Classroom Test</GradingCategory>
      <Description>Pop Quiz</Description>
      <PointsPossible>100</PointsPossible>
      <CreateDate>2000-11-21</CreateDate>
      <DueDate>2000-11-25</DueDate>
      <Weight>1.0</Weight>
      <DetailedDescriptionURL>http://www.assignmentinfo.com/assignment1.pdf</DetailedDescriptionURL> 
    </GradingAssignment>
`
var test_GradingAssignmentScore = `    <GradingAssignmentScore RefId="359D7510-1AD0-A9D7-A8C3-DAD0A85103A2">
      <StudentPersonalRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</StudentPersonalRefId>
      <StudentPersonalLocalId>fred12</StudentPersonalLocalId>
      <TeachingGroupRefId>D0A0A27A-D0A8-510A-D9D7-5101A8C3DA39</TeachingGroupRefId>
      <SchoolInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</SchoolInfoRefId>
      <GradingAssignmentRefId>359D7510-1AD0-A9D7-A8C3-DAD0A85105D2</GradingAssignmentRefId>
      <ScorePoints>45</ScorePoints> 
    </GradingAssignmentScore>
`
var test_example0104 = `    <Identity RefId="4286194F-43ED-43C1-8EE2-F0A27C4BEF86">
      <SIF_RefId SIF_RefObject="StudentPersonal">23B08571-E4D6-45C3-B82A-3E52E5349925</SIF_RefId>
      <AuthenticationSource>AUAccessShibboleth</AuthenticationSource>
      <IdentityAssertions>
        <IdentityAssertion SchemaName="eduPersonPrincipalName">john.doe@asdf.edu.au</IdentityAssertion>
      </IdentityAssertions>
      <AuthenticationSourceGlobalUID>A9A6CB2B-C493-4427-8C1F-D6587D448B35</AuthenticationSourceGlobalUID>
    </Identity>
`
var test_IdentitypublishedbyMicrosoftActiveDirectory = `    <Identity RefId="4286194F-43ED-43C1-8EE2-F0A27C4BEF86">
      <SIF_RefId SIF_RefObject="StudentPersonal">23B08571-E4D6-45C3-B82A-3E52E5349925</SIF_RefId>
      <AuthenticationSource>MSActiveDirectory</AuthenticationSource>
      <IdentityAssertions>
        <IdentityAssertion SchemaName="sAmAccountName">user01</IdentityAssertion>
        <IdentityAssertion SchemaName="userPrincipalName">user01@asdf.edu.au</IdentityAssertion>
        <IdentityAssertion SchemaName="distinguishedName">cn=User01,cn=Users,dc=org</IdentityAssertion>
      </IdentityAssertions>
      <PasswordList>
		  <Password Algorithm="SHA1" KeyName=" ">UGFzc3cwcmQ=</Password>
      </PasswordList>
      <AuthenticationSourceGlobalUID>23A08571-E4D6-45C3-B82A-3E52E5349925</AuthenticationSourceGlobalUID>
    
    <LocalCodeList>
		<LocalCode>
		   <LocalisedCode>SHA-256</LocalisedCode>
		   <Description>This is the algorithm to use here.</Description>
		   <Element> PasswordList/Password[1]/@Algorithm</Element>
		   <ListIndex>1</ListIndex> 
    </LocalCode>
    </LocalCodeList>
    </Identity>
`
var test_IdentitypublishedbyanOpenIDprovider = `  <Identity RefId="4286194F-43ED-43C1-8EE2-F0A27C4BEF86">
    <SIF_RefId SIF_RefObject="StudentPersonal">23B08571-E4D6-45C3-B82A-3E52E5349925</SIF_RefId>
    <AuthenticationSource>OpenID</AuthenticationSource>
    <IdentityAssertions>
      <IdentityAssertion SchemaName="openid.identity">http://verisign.org/p/alice </IdentityAssertion>
      <IdentityAssertion SchemaName="openid.server">http://verisign.org </IdentityAssertion>
    </IdentityAssertions>
    <AuthenticationSourceGlobalUID>A9A6CB2B-C493-4427-8C1F-D6587D448B35</AuthenticationSourceGlobalUID>
  </Identity>
`
var test_LEAInfo = `    <LEAInfo RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      <LocalId>EMR</LocalId>
      <StateProvinceId>EMR</StateProvinceId>
      <CommonwealthId>4215750</CommonwealthId>
      <LEAName>Eastern Metro School District</LEAName>
      <LEAURL>http://www.EMR.vic.edu.au</LEAURL>
      <EducationAgencyType>
        <Code>01</Code>
      </EducationAgencyType>
      <LEAContactList>
        <LEAContact>
          <PublishInDirectory>Y</PublishInDirectory>
          <ContactInfo>
            <Name Type="LGL">
              <FamilyName>Smith</FamilyName>
              <GivenName>Theo</GivenName>
              <FullName>Theo Smith</FullName>
            </Name>
            <PositionTitle>Superintendent</PositionTitle>
            <EmailList>
              <Email Type="01">drseuss@edumail.vic.edu.au</Email>
            </EmailList>
            <PhoneNumberList>
              <PhoneNumber Type="0096">
                <Number>(03) 9745 7455</Number>
              </PhoneNumber>
            </PhoneNumberList>
          </ContactInfo>
        </LEAContact>
      </LEAContactList>
      <PhoneNumberList>
        <PhoneNumber Type="0096">
          <Number>(03) 9745 7455</Number>
        </PhoneNumber>
      </PhoneNumberList>
      <OperationalStatus>O</OperationalStatus>
      <JurisdictionLowerHouse>Unknown</JurisdictionLowerHouse>
      <SLA>205801452</SLA>
    </LEAInfo>

`
var test_LearningResource = `    <LearningResource RefId="6938D88F-82A3-7EFD-8889-2429EC6AB1B5">
      
      <Name>A World of Music (Teacher Reference Sheet)</Name>
      <Author>The Queensland Performing Arts Edge</Author>
      <Contacts>
        <Contact>
          <Address Type="0123" Role="012A">
            <Street>
              <Line1>QPAC</Line1>
              <Line2>Southbank</Line2>
            </Street>
            <City>Brisbane</City>
            <StateProvince>QLD</StateProvince>
            <Country>1101</Country>
            <PostalCode>4002</PostalCode>
          </Address>
          <PhoneNumber Type="0096">
            <Number>(555) 555-1234</Number>
            <ListedStatus>Y</ListedStatus>
          </PhoneNumber>
          <Email Type="01">info@QPAC.edu.au</Email>
        </Contact>
      </Contacts>
      <Location ReferenceType="URI">http://www.artsedge/QPAC/curric/86_u_dreams/489_words.pdf</Location>
      <Status>Available</Status>
      <Description>Teacher Reference Sheet</Description>
      <YearLevels>
        <YearLevel>
          <Code>3</Code>
        </YearLevel>
        <YearLevel>
          <Code>4</Code>
        </YearLevel>
        <YearLevel>
          <Code>5</Code>
        </YearLevel>
      </YearLevels>
      <SubjectAreas>
        <ACStrandSubjectArea>
       <ACStrand>H</ACStrand>
      <SubjectArea>
      	<Code>Hist</Code>
      	  <OtherCodeList>
        	<OtherCode Codeset="Text">Revolutions</OtherCode>
      	</OtherCodeList>
      </SubjectArea>
    </ACStrandSubjectArea>
      </SubjectAreas>
      <MediaTypes>
        <MediaType>x-application/pdf</MediaType>
      </MediaTypes>
      <UseAgreement>see http://www.artsedge.QPAC/teaching_materials/using/artsedge.html for the user agreement</UseAgreement>
      <AgreementDate>2002-06-05</AgreementDate>
      <Approvals>
        <Approval>
          <Organization>QLD DET</Organization>
          <Date>2002-03-15</Date>
        </Approval>
      </Approvals>
      <Evaluations>
        <Evaluation RefId="F1BC63DF-D02C-CED6-54EF-558E84904E01">
          <Description>This resource references QLD curriculum standards</Description>
          <Date>2001-06-15</Date>
          <Name Type="LGL">
            <FamilyName>Doe</FamilyName>
            <GivenName>Joan</GivenName>
          </Name>
        </Evaluation>
      </Evaluations>
      <Components>
        <Component>
          <Name>Vocabulary</Name>
          <Reference>http://www.artsedge.QPAC.org/teaching_materials/curricula/curric/86_u_dreams/489_vocab.html</Reference>
          <Description>Students learn about social contexts such as nationality, culture, etc.</Description>
          <Strategies>
            <Strategy>Inquiry</Strategy>
          </Strategies>
          <AssociatedObjects>
            <AssociatedObject SIF_RefObject="LearningResource">667A87E2-1B4B-9470-CE61-568576DF921E</AssociatedObject>
          </AssociatedObjects>
        </Component>
      </Components>
      <LearningStandards>
        <LearningStandardItemRefId>D56D76D4-0F0B-9691-6DA4-CA2E230494E8</LearningStandardItemRefId>
        <LearningStandardItemRefId>DBCE6406-0B47-E555-A64A-F7FCE6C1A4A9</LearningStandardItemRefId>
      </LearningStandards>
      <LearningResourcePackageRefId>6D254047-C5E5-42CB-B792-1D03ED3BD894</LearningResourcePackageRefId>
    </LearningResource>
`
var test_LearningResourcePackage = `		<LearningResourcePackage RefId="C7DE8668-5968-459F-BF9F-ED22A0E1EA6E">
			<BinaryData MIMEType="text/html;charset=utf-16">0000 0000 0010 0100</BinaryData>
		</LearningResourcePackage>
`
var test_LearningStandardDocument = `       
    <LearningStandardDocument RefId="A5A575C7-8917-5101-B8E7-F08ED123A823">
    <Title>Australian National Curriculum: History</Title>
    <Description>
            AIMS: The aim of the History curriculum is to instil in students, an appreciation in the diversity of the past and how they inform current developments etc etc.
RATIONALE: History is essential for an informed citizenry and has always been a cornerstone of education.
    </Description>
  <Source>National</Source>
  <Organizations>
    <Organization>ACARA</Organization>
  </Organizations>
  <Authors>
    <Author>Joe Bloggs</Author>
  </Authors>
  <OrganizationContactPoint>http://www.acara.edu.au</OrganizationContactPoint>
  <SubjectAreas>
    <ACStrandSubjectArea>
       <ACStrand>H</ACStrand>
      <SubjectArea>
      	<Code>Hist</Code>
      	  <OtherCodeList>
        	<OtherCode Codeset="Text">History</OtherCode>
      	</OtherCodeList>
      </SubjectArea>
    </ACStrandSubjectArea>
  </SubjectAreas>
  <DocumentStatus>Adopted</DocumentStatus>
  <DocumentDate>2010-04-15</DocumentDate>
  <LocalAdoptionDate>2011-01-06</LocalAdoptionDate>
  <EndOfLifeDate>2011-04-15</EndOfLifeDate>
  <Copyright>
    <Date>2010-06-05</Date>
    <Holder>ACARA</Holder>
  </Copyright>
  <YearLevels>
    <YearLevel>
      <Code>11</Code>
    </YearLevel>
    <YearLevel>
      <Code>12</Code>
    </YearLevel>
</YearLevels>
  <RepositoryDate>2011-04-15</RepositoryDate>
  <LearningStandardItemRefId>667A87E2-1B4B-9470-CE61-568576DF921E</LearningStandardItemRefId>
  <RelatedLearningStandards>
    <LearningStandardDocumentRefId>B87A87E2-1B4B-9470-CE61-568576DF921E</LearningStandardDocumentRefId>
  </RelatedLearningStandards>
</LearningStandardDocument>

`
var test_LearningStandardItem = `    <LearningStandardItem RefId="A5D75F78-9175-101B-8C7E-08EA123A8234">
    <StandardSettingBody>
    <Country>1101</Country>
    <StateProvince>VIC</StateProvince>
    <SettingBodyName>DEECD</SettingBodyName>
  </StandardSettingBody>
  <StandardHierarchyLevel>
    <Number>3</Number>
    <Description>AchievementStandard</Description>
  </StandardHierarchyLevel>
  <PredecessorItems>
    <LearningStandardItemRefId>DE072A87-EFAD-4B77-8AF9-FAF83C94839E</LearningStandardItemRefId>
  </PredecessorItems>
  <StatementCodes>
    <StatementCode>Mathematics.2.03.a</StatementCode>
  </StatementCodes>
  <Statements>
    <Statement>Counting from 0 to 100 using whole numbers</Statement>
  </Statements>
  <YearLevels>
    <YearLevel>
      <Code>4</Code>
    </YearLevel>
    <YearLevel>
      <Code>5</Code>
    </YearLevel>
  </YearLevels>
  <ACStrandSubjectArea>
    <ACStrand>M</ACStrand>
    <SubjectArea>  
    	<Code>Maths</Code>
    	<OtherCodeList>
      		<OtherCode Codeset="Text">Mathematics</OtherCode>
    	</OtherCodeList>
    </SubjectArea>
  </ACStrandSubjectArea>


  <StandardIdentifier>
    <YearCreated>2003</YearCreated>
    <ACStrandSubjectArea>
       <ACStrand>M</ACStrand>
      <SubjectArea>
      	<Code>Maths</Code>
      	  <OtherCodeList>
        	<OtherCode Codeset="Text">Mathematics</OtherCode>
      	</OtherCodeList>
      </SubjectArea>
    </ACStrandSubjectArea>
    <StandardNumber>04</StandardNumber>
    <YearLevels>
      <YearLevel>
        <Code>5</Code>
      </YearLevel>
      <YearLevel>
        <Code>6</Code>
      </YearLevel>
      <YearLevel>
        <Code>7</Code>
      </YearLevel>
    </YearLevels>
<Organization>DEECD VIC</Organization>
  </StandardIdentifier>
  <LearningStandardDocumentRefId>8454189F-6BBC-26C6-B97D-DB4B6D0E3AC8</LearningStandardDocumentRefId>
  <RelatedLearningStandardItems>
    <LearningStandardItemRefId RelationshipType="Content">869670E8-540A-4350-9515-AFB767FADD9A</LearningStandardItemRefId>
    <LearningStandardItemRefId RelationshipType="State">B14FDAB4-37CA-4565-A976-5A5B824545C4</LearningStandardItemRefId>
    <LearningStandardItemRefId RelationshipType="ContentElaboration">D5440A11-F5A4-4AC1-9202-61ECBE5A29F6</LearningStandardItemRefId>
  </RelatedLearningStandardItems>
</LearningStandardItem>

`
var test_MarkValueInfo = `		<MarkValueInfo RefId="11737E21-4A7C-46BD-BA43-01CADCA75C87" >
		  <SchoolInfoRefId>A137D78A-E00B-C744-EF90-F2871CEB90A2</SchoolInfoRefId>
		  <Name>Letter Grades</Name>
		  <ValidLetterMarkList>
			<ValidLetterMark>
			  <Code>A</Code>
			  <NumericEquivalent>100</NumericEquivalent>
			</ValidLetterMark>
			<ValidLetterMark>
			  <Code>B</Code>
			  <NumericEquivalent>90</NumericEquivalent>
			</ValidLetterMark>
			<ValidLetterMark>
			  <Code>C</Code>
			  <NumericEquivalent>80</NumericEquivalent>
			</ValidLetterMark>
			<ValidLetterMark>
			  <Code>D</Code>
			  <NumericEquivalent>70</NumericEquivalent>
			</ValidLetterMark>
			<ValidLetterMark>
			  <Code>F</Code>
			  <NumericEquivalent>60</NumericEquivalent>
			</ValidLetterMark>
		 </ValidLetterMarkList>
	     <Narrative>A Narrative about this Grade, letters and their numeric equivalent</Narrative>
	  </MarkValueInfo>
`
var test_PersonPicture = `    <PersonPicture RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      <ParentObjectRefId SIF_RefObject="StudentPersonal">AA648462-8886-24AA-5294-BC638017320B</ParentObjectRefId>
      <SchoolYear>2007</SchoolYear>
      <PictureSource Type="01">http://www.schoolsite.com/pictures/2007/1234.jpg </PictureSource>
      <OKToPublish>Y</OKToPublish>
        <PublishingPermissionList>
           <PublishingPermission>
               <PermissionCategory>OKPrintedMaterial</PermissionCategory>       
               <PermissionValue>Y</PermissionValue>
         </PublishingPermission>
         <PublishingPermission>
               <PermissionCategory>OKOnlineMaterial</PermissionCategory> 
               <PermissionValue>Y</PermissionValue>
         </PublishingPermission>        
         <PublishingPermission>
               <PermissionCategory>OKMediaRelease</PermissionCategory>          
               <PermissionValue>U</PermissionValue>
         </PublishingPermission>               
  </PublishingPermissionList>

    </PersonPicture>

`
var test_PersonPrivacyObligationDocument = `    <PersonPrivacyObligationDocument RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      
      <ParentRefId>AA648462-8886-24AA-5294-BC638017320B</ParentRefId>
      <ParentObjectTypeName>StudentPersonal</ParentObjectTypeName>
      <SchoolYear>2020</SchoolYear>
      <StartDate>2020-02-01</StartDate>
      <EndDate>2020-12-01</EndDate>
      <SettingLocationList>
        <SettingLocation>
          <SettingLocationName>Taylors College</SettingLocationName>
          <SettingLocationType>School</SettingLocationType>
          <SettingLocationRefId>D3E34B35-9D75-101A-8C3D-00AA001A1654</SettingLocationRefId>
          <SettingLocationObjectTypeName>SchoolInfo</SettingLocationObjectTypeName>
        </SettingLocation>
        <SettingLocation>
          <SettingLocationName>Another School</SettingLocationName>
          <SettingLocationType>School</SettingLocationType>
          <SettingLocationRefId>E6E34B35-9D75-101A-8C3D-01AA001A1655</SettingLocationRefId>
          <SettingLocationObjectTypeName>SchoolInfo</SettingLocationObjectTypeName>
        </SettingLocation>
      </SettingLocationList>
      <ContactForRequestsRefId>C4E34B35-9D75-101A-8C3D-00AA001A1653</ContactForRequestsRefId>
      <ContactForRequestsObjectTypeName>StudentContactPersonal</ContactForRequestsObjectTypeName>
      <ConsentToSharingOfData>
        
        
        
        
        <DataDomainObligationList>
          
          <DataDomainObligation>
            <DataDomain>Assessment</DataDomain>
            <DomainComments>Can also share NAPLAN results.</DomainComments>
            <ShareWithList>
              <ShareWith>
                <ShareWithParty>Assessment Vendors</ShareWithParty>
                <ShareWithRefId>AA648462-8886-24AA-5294-BC638017320C</ShareWithRefId>
                <ShareWithObjectTypeName>StudentContactPersonal</ShareWithObjectTypeName>
                <ShareWithLocalId/>
                <ShareWithPurpose>Assessment reporting</ShareWithPurpose>
                <ShareWithRole>Vendor</ShareWithRole>
                <ShareWithComments>A comment</ShareWithComments>
                <PermissionToOnShare>N</PermissionToOnShare>
                <ShareWithURL>www.assessmentworld.com</ShareWithURL>
              </ShareWith>
              <ShareWith>
                <ShareWithParty>Military</ShareWithParty>
                <ShareWithRefId>AA648462-8886-24AA-5294-BC638017320C</ShareWithRefId>
                <ShareWithObjectTypeName>StudentContactPersonal</ShareWithObjectTypeName>
                <ShareWithLocalId/>
                <ShareWithPurpose>entrance requirement</ShareWithPurpose>
                <ShareWithRole>Military</ShareWithRole>
                <ShareWithComments>A comment</ShareWithComments>
                <PermissionToOnShare>N</PermissionToOnShare>
                <ShareWithURL>www.military.com</ShareWithURL>
              </ShareWith>
            </ShareWithList>
          </DataDomainObligation>
          <DataDomainObligation>
            <DataDomain>Attendance</DataDomain>
            <DomainComments>Share with other schools.</DomainComments>
            <ShareWithList>
              <ShareWith>
                <ShareWithParty>Assessment Vendors</ShareWithParty>
                <ShareWithRefId>D3E34B35-9D75-101A-8C3D-00AA001A1651</ShareWithRefId>
                <ShareWithObjectTypeName>StudentContactPersonal</ShareWithObjectTypeName>
                <ShareWithLocalId/>
                <ShareWithPurpose>Assessment reporting</ShareWithPurpose>
                <ShareWithRole>Vendor</ShareWithRole>
                <ShareWithComments>A comment</ShareWithComments>
                <PermissionToOnShare>N</PermissionToOnShare>
                <ShareWithURL>www.assessmentworld.com</ShareWithURL>
              </ShareWith>
              <ShareWith>
                <ShareWithParty>Military</ShareWithParty>
                <ShareWithRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</ShareWithRefId>
                <ShareWithObjectTypeName>StudentContactPersonal</ShareWithObjectTypeName>
                <ShareWithLocalId/>
                <ShareWithPurpose>entrance requirement</ShareWithPurpose>
                <ShareWithRole>Military</ShareWithRole>
                <ShareWithComments>A comment</ShareWithComments> 
                <PermissionToOnShare>N</PermissionToOnShare>
                <ShareWithURL>www.military.com</ShareWithURL>
              </ShareWith>
            </ShareWithList>
            <DoNotShareWithList>
              <DoNotShareWith>
                <DoNotShareWithParty>John Wilbur Loavel</DoNotShareWithParty>
                <DoNotShareWithRefId>B3E34B35-9D75-101A-8C3D-00AA001A1655</DoNotShareWithRefId>
                <DoNotShareWithObjectTypeName>StudentContactPersonal</DoNotShareWithObjectTypeName>
                <DoNotShareWithLocalId/>
                <DoNotShareWithPurpose>Father</DoNotShareWithPurpose>
                <DoNotShareWithRole>Parent</DoNotShareWithRole>
                <DoNotShareWithComments>Student assessment information is not to be shared with the father.</DoNotShareWithComments>
              </DoNotShareWith>
            </DoNotShareWithList>
          </DataDomainObligation>
          
          <DataDomainObligation>
            <DataDomain>Assessment</DataDomain>
            <DomainComments>Can also share NAPLAN results.</DomainComments>
            <ShareWithList>
              <ShareWith>
                <ShareWithParty>Assessment Vendors</ShareWithParty>
                <ShareWithPurpose>Assessment reporting</ShareWithPurpose>
                <ShareWithRole>Assessment</ShareWithRole>
                <ShareWithComments>A comment</ShareWithComments>
                <PermissionToOnShare>N</PermissionToOnShare>
                <ShareWithURL>www.assessmentworld.com</ShareWithURL>
              </ShareWith>
              <ShareWith>
                <ShareWithParty>Military</ShareWithParty>
                <ShareWithPurpose>entrance requirement</ShareWithPurpose>
                <ShareWithRole>Military</ShareWithRole>
                <ShareWithComments>A comment</ShareWithComments>
                <PermissionToOnShare>N</PermissionToOnShare>
                <ShareWithURL>www.military.com</ShareWithURL>
              </ShareWith>
            </ShareWithList>
            <DoNotShareWithList>
              <DoNotShareWith>
                <DoNotShareWithParty>John Wilbur Loavel</DoNotShareWithParty>
                <DoNotShareWithPurpose>Father</DoNotShareWithPurpose>
                <DoNotShareWithRole>Parent</DoNotShareWithRole>
                <DoNotShareWithComments>Student assessment information is not to be shared with the father.</DoNotShareWithComments>
              </DoNotShareWith>
            </DoNotShareWithList>
          </DataDomainObligation>
          <DataDomainObligation>
            <DataDomain>Attendance</DataDomain>
            <DomainComments>Share with other schools.</DomainComments>
            <ShareWithList>
              <ShareWith>
                <ShareWithParty>School XYZ</ShareWithParty>
                <ShareWithPurpose>Attendance reporting</ShareWithPurpose>
                <ShareWithRole>A Role</ShareWithRole>
                <ShareWithComments>A comment</ShareWithComments>
                <PermissionToOnShare>Y</PermissionToOnShare>
                <ShareWithURL>www.XYZ.edu.au</ShareWithURL>
              </ShareWith>
              <ShareWith>
                <ShareWithParty>School ABC</ShareWithParty>
                <ShareWithPurpose>Attendance reporting</ShareWithPurpose>
                <ShareWithRole>A Role</ShareWithRole>
                <ShareWithComments>A comment</ShareWithComments>
                <PermissionToOnShare>Y</PermissionToOnShare>
                <ShareWithURL>www.ABC.edu.au</ShareWithURL>
              </ShareWith>
            </ShareWithList>
            <DoNotShareWithList>
              <DoNotShareWith>
                <DoNotShareWithParty>John Wilbur Loavel</DoNotShareWithParty>
                <DoNotShareWithRefId>F3E34B35-9D75-101A-8C3D-00AA001A1655</DoNotShareWithRefId>
                <DoNotShareWithObjectTypeName>StudentContactPersonal</DoNotShareWithObjectTypeName>
                <DoNotShareWithLocalId/>
                <DoNotShareWithPurpose>Father</DoNotShareWithPurpose>
                <DoNotShareWithRole>Parent</DoNotShareWithRole>
                <DoNotShareWithComments>Student attendance information is not to be shared with the father.</DoNotShareWithComments>
                <DoNotShareWithURL/>
              </DoNotShareWith>
            </DoNotShareWithList>
          </DataDomainObligation>
        </DataDomainObligationList>
        
        <NeverShareWithList>
          <NeverShareWith>
            <NeverShareWithParty>Media</NeverShareWithParty>
            <NeverShareWithRefId>D3E34B35-9D75-101A-8C3D-00AA001A1655</NeverShareWithRefId>
            <NeverShareWithObjectTypeName>StudentContactPersonal</NeverShareWithObjectTypeName>
            <NeverShareWithLocalId/>
            <NeverShareWithPurpose>Never Share</NeverShareWithPurpose>
            <NeverShareWithRole>Company</NeverShareWithRole>
            <NeverShareWithComments>Student and/or family information is not to be shared with the Media.</NeverShareWithComments>
            <NeverShareWithURL>www.Mediaworld.com</NeverShareWithURL>
          </NeverShareWith>
        </NeverShareWithList>
      </ConsentToSharingOfData>
      <PermissionToParticipateList>
        <PermissionToParticipate>
          
          
          
          <PermissionCategory>Travel</PermissionCategory>
          <Permission>WalkHome</Permission>
          <PermissionValue>Y</PermissionValue>
          <PermissionStartDate>2020-01-01</PermissionStartDate>
          <PermissionEndDate>2020-12-31</PermissionEndDate>
          <PermissionGranteeName>Sarah Lain Loavel</PermissionGranteeName>
          <PermissionGranteeRelationship>Mother</PermissionGranteeRelationship>
          <PermissionComments>Only with a sibling.</PermissionComments>
        </PermissionToParticipate>
        <PermissionToParticipate>
          
          
          
          <PermissionCategory>Travel</PermissionCategory>
          <Permission>RideBikeHome</Permission>
          <PermissionValue>N</PermissionValue>
          <PermissionStartDate>2020-01-01</PermissionStartDate>
          <PermissionEndDate>2020-12-31</PermissionEndDate>
          <PermissionComments>Only with a sibling.</PermissionComments>
        </PermissionToParticipate>
        <PermissionToParticipate>
          
          
          
          <PermissionCategory>Excursions</PermissionCategory>
          <Permission>Attend Excursions General</Permission>
          <PermissionValue>Y</PermissionValue>
          <PermissionStartDate>2020-01-01</PermissionStartDate>
          <PermissionEndDate>2020-12-31</PermissionEndDate>
          <PermissionComments>Optional Comment</PermissionComments>
        </PermissionToParticipate>
        <PermissionToParticipate>
          
          
          
          <PermissionCategory>Excursions</PermissionCategory>
          <Permission>Overnight Camps</Permission>
          <PermissionValue>N</PermissionValue>
          <PermissionStartDate>2020-01-01</PermissionStartDate>
          <PermissionEndDate>2020-12-31</PermissionEndDate>
          <PermissionComments>Optional Comment</PermissionComments>
        </PermissionToParticipate>
      </PermissionToParticipateList>
      <ApplicableLawList>
        <ApplicableLaw>
          <ApplicableCountry>AU</ApplicableCountry>
          
          <ApplicableLawName>APP1</ApplicableLawName>
          <ApplicableLawURL>http://www.vic.priv.contract.edu.au</ApplicableLawURL>
        </ApplicableLaw>
        <ApplicableLaw>
          <ApplicableCountry>AU</ApplicableCountry>
          
          <ApplicableLawName>APP2</ApplicableLawName>
          <ApplicableLawURL>http://www.vic.priv.contract.edu.au</ApplicableLawURL>
        </ApplicableLaw>
      </ApplicableLawList>
    </PersonPrivacyObligationDocument>
`
var test_PersonalisedPlan = `    <PersonalisedPlan RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
      <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
      <PersonalisedPlanCategory>M</PersonalisedPlanCategory>
      <PersonalisedPlanStartDate>2017-11-01</PersonalisedPlanStartDate>
      <PersonalisedPlanEndDate>2019-11-10</PersonalisedPlanEndDate>
      <PersonalisedPlanReviewDate>2018-11-10</PersonalisedPlanReviewDate>
      
      <PersonalisedPlanNotes>This Student is allergic to cashews.</PersonalisedPlanNotes>
    </PersonalisedPlan>

`
var test_ResourceBookingExample = `    <ResourceBooking RefId="B5739375-800A-C4CC-6385-0BB2754114AA">
      <ResourceRefId SIF_RefObject="LearningResource">B7A34E56-1C97-345C-0A4E-11BB112B2753</ResourceRefId>
      <ResourceLocalId>ZXY789</ResourceLocalId>
      <StartDateTime>2011-01-01T16:30:00</StartDateTime>  
      <FinishDateTime>2011-01-01T17:30:00</FinishDateTime>
      <FromPeriod>3</FromPeriod>  
      <ToPeriod>3</ToPeriod>  
      <Booker>98157AA0-13BA-8C3D-00AA-012B359D7512</Booker>
      <ScheduledActivityRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</ScheduledActivityRefId>
      <KeepOld>false</KeepOld>
    </ResourceBooking>
    

`
var test_RoomInfoExample = `    <RoomInfo RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652" >
      <SchoolInfoRefId>A2E35B35-9D75-101A-8C3D-00AA001A0000</SchoolInfoRefId>
      <RoomNumber>101</RoomNumber>
      <StaffList>
        <StaffPersonalRefId>A8C3A2E3-5B35-9D75-101D-00AA001A0000</StaffPersonalRefId>
      </StaffList>
      <Description>Room 101</Description>
      <Building>Main A</Building>
      <HomeroomNumber>10-HR-A</HomeroomNumber>
      <Size>400</Size>
      <Capacity>35</Capacity>
      <PhoneNumber Type="0096">
        <Number>(02) 9555-1234</Number>
      </PhoneNumber>
      <AvailableForTimetable>Y</AvailableForTimetable>
    </RoomInfo>
`
var test_ScheduledActivityExample = `      <ScheduledActivity RefId="BD3E34B3-59D7-5101-A8C3-D00AA001A165">
        <SchoolInfoRefId>B7A34E56-1C97-345C-0A4E-11BB112B2753</SchoolInfoRefId>
        <TimeTableCellRefId>FDF05286-9094-4405-96B3-D6CC0ED5F00B</TimeTableCellRefId>
   <DayId>1</DayId>
        <TimeTableRefId>2A7F0D96-3E99-4B83-A958-690C360A36C2</TimeTableRefId>
   <ActivityDate>2014-10-10</ActivityDate>
   <StartTime>12:05:00</StartTime>
   <FinishTime>13:30:00</FinishTime>
   <TeacherList>
     <TeacherCover>
       <StaffPersonalRefId>98157AA0-13BA-8C3D-00AA-012B359D7512</StaffPersonalRefId>
       <StaffLocalId>US8923</StaffLocalId>
       <StartTime>12:05:00</StartTime>
       <FinishTime>13:00:00</FinishTime>
       <Credit>In-Lieu</Credit>
       <Supervision>MinimalSupervision</Supervision>
       <Weighting>0.5</Weighting>
     </TeacherCover>
     <TeacherCover>
       <StaffPersonalRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
       <StaffLocalId>ZY3213</StaffLocalId>
       <StartTime>13:00:00</StartTime>
       <FinishTime>13:30:00</FinishTime>
       <Credit>Extra</Credit>
       <Supervision>MinimalSupervision</Supervision>
       <Weighting>0.5</Weighting>
     </TeacherCover>
   </TeacherList>
   <RoomList>
     <RoomInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</RoomInfoRefId>
   </RoomList>
   <ActivityType>Incursion</ActivityType>
   <StudentList>
     <StudentPersonalRefId>FAE9D90A-38F8-4B72-9C92-F868CB63C9F3</StudentPersonalRefId>
   </StudentList>
   <TeachingGroupList>
     <TeachingGroupRefId>EB79C3D1-FF19-11D7-8513-8B604A511DAD</TeachingGroupRefId>
     <TeachingGroupRefId>B2345163-8474-6B38-7459-000F84723A00</TeachingGroupRefId>
   </TeachingGroupList>
   <Override DateOfOverride="2020-02-02">Y</Override>
   <TimeTableChangeReasonList>
     <TimeTableChangeReason>
	   <TimeTableChangeType>TeacherAbsence</TimeTableChangeType>
	   <TimeTableChangeNotes>SL</TimeTableChangeNotes>
     </TimeTableChangeReason>
     <TimeTableChangeReason>
	   <TimeTableChangeType>RoomRemoval</TimeTableChangeType>
	   <TimeTableChangeNotes>GL</TimeTableChangeNotes>
     </TimeTableChangeReason>
   </TimeTableChangeReasonList>
 </ScheduledActivity> 
`
var test_SchoolCourseInfoExample = `    <SchoolCourseInfo RefId="9D75101A-8C3D-00AA-001A-0000A2E35B35">
      <SchoolInfoRefId>101A8C3D-00AA-001A-0000-A2E35B359D75</SchoolInfoRefId>
      <SchoolYear>2006</SchoolYear>
      <CourseCode>CS101</CourseCode>
      <StateCourseCode>08-001</StateCourseCode>
      <DistrictCourseCode>CS101</DistrictCourseCode>
      <SubjectAreaList>
        <SubjectArea>
          <Code>Graphic Arts</Code>
        </SubjectArea>
      </SubjectAreaList>
      <CourseTitle>Gif, JPeg, or Png: What's the Difference?</CourseTitle>
      <Description>Explore the various types of files related to graphic arts.</Description>
      <InstructionalLevel>0571</InstructionalLevel>
      <CourseCredits>2</CourseCredits>
      <CoreAcademicCourse>N</CoreAcademicCourse>
      <GraduationRequirement>N</GraduationRequirement>
    </SchoolCourseInfo>
`
var test_SchoolInfo = `    <SchoolInfo RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      <LocalId>01011234</LocalId>
      <StateProvinceId>01011234</StateProvinceId>
      <CommonwealthId>012345</CommonwealthId>
      <ParentCommonwealthId>012346</ParentCommonwealthId>
      <SchoolName>Lincoln Secondary College</SchoolName>
      <LEAInfoRefId>73648462-8886-24AA-5294-BC6380173276</LEAInfoRefId>
      <OtherLEA SIF_RefObject="LEAInfo">AA648462-8886-24AA-5294-BC638017320B</OtherLEA>
      <SchoolDistrict>Southern Metropolitan Region</SchoolDistrict>
      <SchoolType>Pri/Sec</SchoolType>
      <SchoolFocusList>
        <SchoolFocus>01</SchoolFocus>
        <SchoolFocus>03</SchoolFocus>
      </SchoolFocusList>
      <SchoolURL>http://www.lincolnsc.edu.vic.au</SchoolURL>
      <PrincipalInfo>
        <ContactName Type="LGL">
            <Title>Mr</Title>
            <FamilyName>Mason</FamilyName>
            <GivenName>Francis</GivenName>
            <MiddleName>Frank</MiddleName>
            <Suffix>Jr.</Suffix>
            <FullName>Mr Frank Mason Jr.</FullName>
        </ContactName>
        <ContactTitle>Senior School Principal</ContactTitle>
      </PrincipalInfo>
      <SchoolContactList>
        <SchoolContact>
          <PublishInDirectory>Y</PublishInDirectory>
          <ContactInfo>
            <Name Type="LGL">
              <Title>Mr</Title>
              <FamilyName>Miller</FamilyName>
              <GivenName>James</GivenName>
              <MiddleName>Mark</MiddleName>
              <Suffix>Jr.</Suffix>
                <FullName>Mr James Mark Miller Jr.</FullName>
              </Name>
            <PositionTitle>Business Manager</PositionTitle>
            <Role>School Information Contact Point</Role>
            <Address Type="0123" Role="012B">
              <Street>
                <Line1>23 Nicholson Street</Line1>
                </Street>
                  <City>Carnegie</City>
                  <StateProvince>VIC</StateProvince>
                  <Country>1101</Country>
                  <PostalCode>3004</PostalCode>
                  <GridLocation>
                    <Latitude>23.9876</Latitude>
                    <Longitude>-98.8765</Longitude>
                  </GridLocation>
                </Address>
            <EmailList>
              <Email Type="01">jmiller@lsc.vic.edu.au</Email>
              <Email Type="02">jmiller@yahoo.com.au</Email>
            </EmailList>
            <PhoneNumberList>
              <PhoneNumber Type="0096">
                <Number>03 9637-2000</Number>
                <Extension>72345</Extension>
                <ListedStatus>Y</ListedStatus>
              </PhoneNumber>
            </PhoneNumberList>
          </ContactInfo>
        </SchoolContact>
      </SchoolContactList>
	  <AddressList>
	             <Address Type="0123" Role="012B">
              <Street>
                <Line1>23 Nicholson Street</Line1>
                </Street>
                  <City>Carnegie</City>
                  <StateProvince>VIC</StateProvince>
                  <Country>1101</Country>
                  <PostalCode>3004</PostalCode>
                  <GridLocation>
                    <Latitude>23.9876</Latitude>
                    <Longitude>-98.8765</Longitude>
                  </GridLocation>
                </Address>
	  </AddressList>
      <PhoneNumberList>
        <PhoneNumber Type="0096">
          <Number>03 9637-2000</Number>
        </PhoneNumber>
      </PhoneNumberList>
      <SessionType>0827</SessionType>
      <YearLevels>
        <YearLevel>
          <Code>6</Code>
        </YearLevel>
        <YearLevel>
          <Code>7</Code>
        </YearLevel>
        <YearLevel>
          <Code>8</Code>
        </YearLevel>
        <YearLevel>
          <Code>9</Code>
        </YearLevel>
        <YearLevel>
          <Code>10</Code>
        </YearLevel>
        <YearLevel>
          <Code>11</Code>
        </YearLevel>
        <YearLevel>
          <Code>12</Code>
        </YearLevel>
      </YearLevels>
      <ARIA>1.0</ARIA>
      <OperationalStatus>O</OperationalStatus>
      <FederalElectorate>216</FederalElectorate>
      <Campus>
        <SchoolCampusId>01</SchoolCampusId>
        <CampusType>Camp</CampusType>
            <AdminStatus>Y</AdminStatus>
          </Campus>
      <SchoolSector>NG</SchoolSector>
      <IndependentSchool>Y</IndependentSchool>
      <NonGovSystemicStatus>S</NonGovSystemicStatus>
      <System>0003</System>
      <ReligiousAffiliation>2171</ReligiousAffiliation>
        <SchoolGeographicLocation>10</SchoolGeographicLocation>
        <LocalGovernmentArea>Cardinia</LocalGovernmentArea>
        <JurisdictionLowerHouse>Unknown</JurisdictionLowerHouse>
        <SLA>205801452</SLA>
        <SchoolCoEdStatus>C</SchoolCoEdStatus>
        <SchoolGroupList>
          <SchoolGroup>YVC</SchoolGroup>
            <SchoolGroup>EastSec01</SchoolGroup>
        </SchoolGroupList>
      </SchoolInfo>
`
var test_SchoolPrograms = `    <SchoolPrograms RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      <SchoolInfoRefId>73648462-8886-24AA-5294-BC6380173276</SchoolInfoRefId>
      <SchoolYear>2009</SchoolYear>
      <SchoolProgramList>
        <Program>
          <Category>01</Category>
          <Type>Steiner program</Type>
        </Program>
        <Program>
          <Category>01</Category>
          <Type>Tournament of minds</Type>
        </Program>
      </SchoolProgramList>
    </SchoolPrograms>

`
var test_SectionInfoExample1 = `		<SectionInfo RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
			<SchoolCourseInfoRefId>101A8C3D-00AA-001A-0000-A2E35B359D75</SchoolCourseInfoRefId>
			<LocalId>287-1</LocalId>
			<SchoolYear>2013</SchoolYear>
			<MediumOfInstruction>
				<Code>0605</Code>
			</MediumOfInstruction>
			<LanguageOfInstruction>
				<Code>1201</Code>
			</LanguageOfInstruction>
			<LocationOfInstruction>
				<Code>0340</Code>
				<OtherCodeList>
					<OtherCode Codeset="Text">NSW DEC</OtherCode>
				</OtherCodeList>
			</LocationOfInstruction>
			<SchoolCourseInfoOverride Override="Yes">
				<CourseCode>CS101A</CourseCode>
				<StateCourseCode>08-001A</StateCourseCode>
				<DistrictCourseCode>CS101A</DistrictCourseCode>
				<SubjectArea>
					<Code>05</Code>
					<OtherCodeList>
						<OtherCode Codeset="Text">Graphic Arts for Beginners</OtherCode>
					</OtherCodeList>
				</SubjectArea>
				<CourseTitle>Graphics Basics</CourseTitle>
				<InstructionalLevel>Graduate Certificate II</InstructionalLevel>
				
			</SchoolCourseInfoOverride>
		</SectionInfo>
`
var test_SessionInfoExample = `    <SessionInfo RefId="98157AA0-13BA-8C3D-00AA-012B359D7512">
      <SchoolInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</SchoolInfoRefId>
      <TimeTableCellRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</TimeTableCellRefId>
      <SchoolYear>2008</SchoolYear>
      <LocalId>2</LocalId>
      <TimeTableSubjectLocalId>10MA1</TimeTableSubjectLocalId>
      <TeachingGroupLocalId>10MA1-Smith</TeachingGroupLocalId>
      <SchoolLocalId>01991</SchoolLocalId>
      <StaffPersonalLocalId>SMI009</StaffPersonalLocalId>
      <RoomNumber>R08</RoomNumber>
      <DayId>1</DayId>
          <PeriodId>5</PeriodId>
          <SessionDate>2016-10-10</SessionDate>
          <StartTime>12:05:00</StartTime>
          <FinishTime>13:30:00</FinishTime>
          <RollMarked>Y</RollMarked>
        </SessionInfo>
`
var test_StaffAssignment = `    <StaffAssignment RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      <SchoolInfoRefId>A8C3D3E3-4B35-9D75-101D-00AA001A1652</SchoolInfoRefId>
      <SchoolYear>2008</SchoolYear>
      <StaffPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1651</StaffPersonalRefId>
      <Description>VCE English Teacher</Description>
      <PrimaryAssignment>Y</PrimaryAssignment>
      <JobStartDate>2000-09-05</JobStartDate>
      <JobEndDate>2001-06-25</JobEndDate>
      <JobFTE>1.00</JobFTE>
      <JobFunction>Teacher</JobFunction>
      <StaffSubjectList>
        <StaffSubject>
          <PreferenceNumber>1</PreferenceNumber>
          <SubjectLocalId>English</SubjectLocalId>
        </StaffSubject>
        <StaffSubject>
          <PreferenceNumber>2</PreferenceNumber>
          <SubjectLocalId>Mathematics</SubjectLocalId>
        </StaffSubject>
        <StaffSubject>
          <PreferenceNumber>3</PreferenceNumber>
          <SubjectLocalId>Science</SubjectLocalId>
        </StaffSubject>
      </StaffSubjectList>
      <YearLevels>
        <YearLevel>
          <Code>11</Code>
        </YearLevel>
        <YearLevel>
          <Code>12</Code>
        </YearLevel>
      </YearLevels>
      <CasualReliefTeacher>N</CasualReliefTeacher>
      <AvailableForTimetable>Y</AvailableForTimetable>
    </StaffAssignment>

`
var test_StaffPersonal = `    <StaffPersonal RefId="D3E34F41-9D75-101A-8C3D-00AA001A1652">
      <LocalId>946379881</LocalId>
      <StateProvinceId>C2345681</StateProvinceId>
      <ElectronicIdList>
        <ElectronicId Type="01">206655</ElectronicId>
      </ElectronicIdList>
      <OtherIdList>
        <OtherId Type="0004">333333333</OtherId>
      </OtherIdList>
      <PersonInfo>
        <Name Type = "LGL" >
          <FamilyName>Smith</FamilyName>
          <GivenName>Fred</GivenName>
          <FullName>Fred Smith</FullName>
        </Name>
        <OtherNames>
          <Name Type="AKA">
            <FamilyName>Anderson</FamilyName>
            <GivenName>Samuel</GivenName>
            <FullName>Samuel Anderson</FullName>
          </Name>
          <Name Type="PRF">
            <FamilyName>Rowinski</FamilyName>
            <GivenName>Sam</GivenName>
            <FullName>Sam Rowinski </FullName>
          </Name>
        </OtherNames>

        <Demographics>
          <IndigenousStatus>3</IndigenousStatus>
          <Sex>1</Sex>
          <BirthDate>1990-09-26</BirthDate>
          <BirthDateVerification>1004</BirthDateVerification>
          <PlaceOfBirth>Clayton</PlaceOfBirth>
          <StateOfBirth>VIC</StateOfBirth>
          <CountryOfBirth>1101</CountryOfBirth>
          <CountriesOfCitizenship>
            <CountryOfCitizenship>8104</CountryOfCitizenship>
            <CountryOfCitizenship>1101</CountryOfCitizenship>
          </CountriesOfCitizenship>
          <CountriesOfResidency>
            <CountryOfResidency>8104</CountryOfResidency>
            <CountryOfResidency>1101</CountryOfResidency>
          </CountriesOfResidency>
          <CountryArrivalDate>1990-09-26</CountryArrivalDate>
          <AustralianCitizenshipStatus>1</AustralianCitizenshipStatus>
          <EnglishProficiency>
            <Code>1</Code>
          </EnglishProficiency>
          <LanguageList>
            <Language>
              <Code>0001</Code>
              <LanguageType>1</LanguageType>
            </Language>
          </LanguageList>
          <DwellingArrangement>
            <Code>1671</Code>
          </DwellingArrangement>
          <Religion>
            <Code>2013</Code>
          </Religion>
          <ReligiousEventList>
            <ReligiousEvent>
              <Type>Baptism</Type>
              <Date>2000-09-01</Date>
            </ReligiousEvent>
            <ReligiousEvent>
              <Type>Christmas</Type>
              <Date>2009-12-24</Date>
            </ReligiousEvent>
          </ReligiousEventList>
          <ReligiousRegion>The Religion Region</ReligiousRegion>
          <PermanentResident>P</PermanentResident>
          <VisaSubClass>101</VisaSubClass>
          <VisaStatisticalCode>05</VisaStatisticalCode>
        </Demographics>
        <AddressList>
          <Address Type ="0123" Role="012A">
            <Street>
              <Line1>Unit1/10</Line1>
              <Line2>Barkley Street</Line2>
            </Street>
            <City>Yarra Glenn</City>
            <StateProvince>VIC</StateProvince>
            <Country>1101</Country>
            <PostalCode>9999</PostalCode>
          </Address>
          <Address Type="0123A" Role="1073">
            <Street>
              <Line1>34 Term Address Street</Line1>
            </Street>
            <City>Home Town</City>
            <StateProvince>WA</StateProvince>
            <Country>1101</Country>
            <PostalCode>9999</PostalCode>
          </Address>
        </AddressList>
        <PhoneNumberList>
          <PhoneNumber Type="0096">
            <Number>03 9637-2289</Number>
            <Extension>72289</Extension>
            <ListedStatus>Y</ListedStatus>
          </PhoneNumber>
          <PhoneNumber Type="0888">
            <Number>0437-765-234</Number>
            <ListedStatus>N</ListedStatus>
          </PhoneNumber>
        </PhoneNumberList>
        <EmailList>
          <Email Type="01">fsmith@yahoo.com</Email>
          <Email Type="02">freddy@gmail.com</Email>
        </EmailList>
      </PersonInfo>

      <Title>Principal</Title>
      <EmploymentStatus>A</EmploymentStatus>
      
       <MostRecent>
         <SchoolLocalId>S12345</SchoolLocalId>
		 <LocalCampusId>D</LocalCampusId>
		 
		 <NAPLANClassList>
				<ClassCode>9</ClassCode>
				<ClassCode>7</ClassCode>
		 </NAPLANClassList>
         <HomeGroup>9G</HomeGroup>
	   </MostRecent>

      
    </StaffPersonal>
   
`
var test_StudentActivityInfo = `    <StudentActivityInfo RefId="6472B261-0947-583A-463D-BB345291B001">
      <Title>Book Club</Title>
      <Description>Group of middle school students promoting reading</Description>
      <StudentActivityType>
        <Code>6011</Code>
        <OtherCodeList>
          <OtherCode Codeset="Local">MBook</OtherCode>
        </OtherCodeList>
      </StudentActivityType>
      <StudentActivityLevel>Middle School</StudentActivityLevel>
      <YearLevels>
        <YearLevel>
          <Code>5</Code>
        </YearLevel>
        <YearLevel>
          <Code>6</Code>
        </YearLevel>
        <YearLevel>
          <Code>7</Code>
        </YearLevel>
        <YearLevel>
          <Code>8</Code>
        </YearLevel>
      </YearLevels>
      <CurricularStatus>0750</CurricularStatus>
      <Location Type="Classroom">
        <LocationName>Beaconhills Middle School Library</LocationName>
        <LocationRefId SIF_RefObject="RoomInfo">94758261-0947-583A-CEB2-BB345291BAAA</LocationRefId>
      </Location>
    </StudentActivityInfo>

`
var test_StudentActivityParticipation = `    <StudentActivityParticipation RefId="9ECC9683-0E02-406F-926C-1C4D3542D122">
      <StudentPersonalRefId>646C5D4A-C829-4886-A02B-971695C7BC06</StudentPersonalRefId>
      <StudentActivityInfoRefId>6472B261-0947-583A-463D-BB345291B001</StudentActivityInfoRefId>
      <SchoolYear>2009</SchoolYear>
      <ParticipationComment>First year of participation</ParticipationComment>
      <StartDate>2008-09-01</StartDate>
      <Role>Team member</Role>
      <RecognitionList>
        <Recognition>0750</Recognition>
      </RecognitionList>
    </StudentActivityParticipation>

`
var test_StudentAttendanceSummary = `    <StudentAttendanceSummary
        StudentAttendanceSummaryRefId="D3476FAE-8647-384B-DA24-31EDA3583211" >
        <StudentPersonalRefId>7C834EA9-EDA1-2090-347F-83297E1C290C</StudentPersonalRefId>
        <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
        <SchoolYear>2015</SchoolYear>
        <StartDate>2014-08-30</StartDate>
        <EndDate>2015-06-10</EndDate>  
      
      <StartDay>1</StartDay>
      
      <EndDay>180</EndDay>
      <FTE>1.00</FTE>
      <DaysAttended>178</DaysAttended>
      <ExcusedAbsences>2</ExcusedAbsences>
      <UnexcusedAbsences>0</UnexcusedAbsences>
      <DaysTardy>3</DaysTardy>
      <DaysInMembership>180</DaysInMembership>
    </StudentAttendanceSummary>
`
var test_StudentAttendanceTimeListExample = `		<StudentAttendanceTimeList RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
			<StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
			<SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
			<Date>2002-11-01</Date>
			<SchoolYear>2003</SchoolYear>
			<AttendanceTimes>

				<AttendanceTime>
					<AttendanceCode>
						<Code>100</Code>
					</AttendanceCode>
					<AttendanceStatus>01</AttendanceStatus>
					<StartTime>09:00:00</StartTime>
					<EndTime>11:00:00</EndTime>
					<AttendanceNote/>
				</AttendanceTime>
				<AttendanceTime>
					<AttendanceCode>
						<Code>200</Code>
						<OtherCodeList>
							<OtherCode Codeset="Local">S</OtherCode>
							<OtherCode Codeset="Text">C</OtherCode>
						</OtherCodeList>
					</AttendanceCode>
					<AttendanceStatus>01</AttendanceStatus>
					<StartTime>11:00:00</StartTime>
					<EndTime>12:05:00</EndTime>
					<DurationValue>0.167</DurationValue>
					<AttendanceNote>Left for Orthodontist appt. and returned to school
						afterward</AttendanceNote>
				</AttendanceTime>
				<AttendanceTime>
					<AttendanceCode>
						<Code>100</Code>
					</AttendanceCode>
					<AttendanceStatus>01</AttendanceStatus>
					<StartTime>12:05:00</StartTime>
					<EndTime>15:30:00</EndTime>
					<AttendanceNote/>
				</AttendanceTime>
			</AttendanceTimes>
		</StudentAttendanceTimeList>
`
var test_StudentAttendanceTimeListExample2 = `    <StudentAttendanceTimeList RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
    <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
	<SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
	<Date>2002-11-01</Date>
	<SchoolYear>2003</SchoolYear>
	<AttendanceTimes>

			<AttendanceTime>
				<AttendanceType>LA</AttendanceType> 
				<AttendanceCode>
					<Code>102</Code>
				</AttendanceCode>
				<AttendanceStatus>02</AttendanceStatus>
				<StartTime>09:00:00</StartTime>
				<EndTime>09:30:00</EndTime>
				<AttendanceNote>No acceptable excuse given</AttendanceNote>
			</AttendanceTime>
			<AttendanceTime>
				<AttendanceType>RC</AttendanceType>	 
				<AttendanceCode>
					<Code>100</Code>
				</AttendanceCode>
				<AttendanceStatus>NA</AttendanceStatus>
				<StartTime>09:30:00</StartTime>
				<EndTime>11:00:00</EndTime>
				<DurationValue>1.000</DurationValue>
			</AttendanceTime>
			<AttendanceTime>
				<AttendanceType>OC</AttendanceType> 
				<AttendanceCode>
					<Code>200</Code>
				</AttendanceCode>
				<AttendanceStatus>01</AttendanceStatus>
				<StartTime>11:00:00</StartTime>
				<EndTime>14:30:00</EndTime>
				<AttendanceNote>Student had stomach ache</AttendanceNote>
			</AttendanceTime>
			<AttendanceTime>
				<AttendanceType>ED</AttendanceType> 
				<AttendanceCode>
					<Code>205</Code>
				</AttendanceCode>
				<AttendanceStatus>01</AttendanceStatus>
				<StartTime>14:30:00</StartTime>
				<EndTime>15:00:00</EndTime>
				<AttendanceNote>Seeing the doctor in regards to stomach ache</AttendanceNote>
			</AttendanceTime>
	</AttendanceTimes>
	<PeriodAttendances>
			<PeriodAttendance>
				<AttendanceType>CR</AttendanceType> 
				<AttendanceCode>
					<Code>102</Code>
				</AttendanceCode>
				<AttendanceStatus>02</AttendanceStatus>
				<Date>2002-11-01</Date> 
				<TimetablePeriod>P1</TimetablePeriod>
				<DayId>1</DayId>
				<StartTime>09:00:00</StartTime>
				<EndTime>10:00:00</EndTime>
				<TimeTableSubjectRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</TimeTableSubjectRefId>
				<TeacherList>
					<TeacherCover>
						<StaffPersonalRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
					</TeacherCover>
					<TeacherCover>
						<StaffPersonalRefId>BB5A0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
					</TeacherCover>
				</TeacherList>
				<RoomList>
					<RoomInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</RoomInfoRefId>
				</RoomList>
			</PeriodAttendance>
			<PeriodAttendance>
				<AttendanceType>CR</AttendanceType> 
				<AttendanceCode>
					<Code>100</Code>
				</AttendanceCode>
				<AttendanceStatus>NA</AttendanceStatus>
				<Date>2002-11-01</Date>
				<TimetablePeriod>P2</TimetablePeriod>
				<DayId>1</DayId>
				<StartTime>10:00:00</StartTime>
				<EndTime>11:00:00</EndTime>
				<TimeTableSubjectRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</TimeTableSubjectRefId>
				<TeacherList>
					<TeacherCover>
						<StaffPersonalRefId>AB5A0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
					</TeacherCover>
				</TeacherList>
				<RoomList>
					<RoomInfoRefId>11237EA4-301C-ADCA-75C8-7214A7C46BDB</RoomInfoRefId>
				</RoomList>
			</PeriodAttendance>
			<PeriodAttendance>
				<AttendanceType>CR</AttendanceType> 
				<AttendanceCode>
					<Code>101</Code>
				</AttendanceCode>
				<AttendanceStatus>01</AttendanceStatus>
				<Date>2002-11-01</Date>				
				<TimetablePeriod>P3</TimetablePeriod>
				<DayId>1</DayId>
				<StartTime>11:00:00</StartTime>
				<EndTime>12:00:00</EndTime>
				<TimeTableSubjectRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</TimeTableSubjectRefId>
				<TeacherList>
					<TeacherCover>
						<StaffPersonalRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
					</TeacherCover>
					<TeacherCover>
						<StaffPersonalRefId>A7BB0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
					</TeacherCover>
				</TeacherList>
				<RoomList>
					<RoomInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</RoomInfoRefId>
				</RoomList>
			</PeriodAttendance>
			<PeriodAttendance>
				<AttendanceType>CR</AttendanceType>
				<AttendanceCode>
					<Code>100</Code>
				</AttendanceCode>
				<AttendanceStatus>NA</AttendanceStatus>
				<Date>2002-11-01</Date>
				<TimetablePeriod>P4</TimetablePeriod>
				<DayId>1</DayId>
				<StartTime>13:00:00</StartTime>
				<EndTime>14:00:00</EndTime>
				<TimeTableSubjectRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</TimeTableSubjectRefId>
				<TeacherList>
					<TeacherCover>
						<StaffPersonalRefId>A75F0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
					</TeacherCover>
				</TeacherList>
				<RoomList>
					<RoomInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</RoomInfoRefId>
				</RoomList>
			</PeriodAttendance>
			<PeriodAttendance>
				<AttendanceType>CR</AttendanceType> 
				<AttendanceCode>
					<Code>100</Code>
				</AttendanceCode>
				<AttendanceStatus>NA</AttendanceStatus>
				<Date>2002-11-01</Date>
				<TimetablePeriod>P5</TimetablePeriod>
				<DayId>1</DayId>
				<StartTime>14:00:00</StartTime>
				<EndTime>15:00:00</EndTime>
				<TimeTableSubjectRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</TimeTableSubjectRefId>
				<TeacherList>
					<TeacherCover>
						<StaffPersonalRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
					</TeacherCover>
				</TeacherList>
				<RoomList>
					<RoomInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</RoomInfoRefId>
				</RoomList>
			</PeriodAttendance>
		</PeriodAttendances>
</StudentAttendanceTimeList>
`
var test_StudentContactPersonalExample = `    <StudentContactPersonal RefId='7C834EA9-EDA1-2090-347F-83297E1C290F'>
      <LocalId>P1234567</LocalId>
      <OtherIdList>
        <OtherId Type="Health Care Card">098765</OtherId>
        <OtherId Type="Seniors Card">123456789</OtherId>
      </OtherIdList>
      <PersonInfo>
        <Name Type = "LGL" >
          <FamilyName>Smith</FamilyName>
          <GivenName>Fred</GivenName>
          <FullName>Fred Smith</FullName>
        </Name>
        <OtherNames>
          <Name Type="AKA">
            <FamilyName>Anderson</FamilyName>
            <GivenName>Samuel</GivenName>
            <FullName>Samuel Anderson</FullName>
          </Name>
          <Name Type="PRF">
            <FamilyName>Rowinski</FamilyName>
            <GivenName>Sam</GivenName>
            <FullName>Sam Rowinski </FullName>
          </Name>
        </OtherNames>

        <Demographics>
          <IndigenousStatus>3</IndigenousStatus>
          <Sex>1</Sex>
          <BirthDate>1990-09-26</BirthDate>
          <BirthDateVerification>1004</BirthDateVerification>
          <PlaceOfBirth>Clayton</PlaceOfBirth>
          <StateOfBirth>VIC</StateOfBirth>
          <CountryOfBirth>1101</CountryOfBirth>
          <CountriesOfCitizenship>
            <CountryOfCitizenship>8104</CountryOfCitizenship>
            <CountryOfCitizenship>1101</CountryOfCitizenship>
          </CountriesOfCitizenship>
          <CountriesOfResidency>
            <CountryOfResidency>8104</CountryOfResidency>
            <CountryOfResidency>1101</CountryOfResidency>
          </CountriesOfResidency>
          <CountryArrivalDate>1990-09-26</CountryArrivalDate>
          <AustralianCitizenshipStatus>1</AustralianCitizenshipStatus>
          <EnglishProficiency>
            <Code>1</Code>
          </EnglishProficiency>
          <LanguageList>
            <Language>
              <Code>0001</Code>
              <LanguageType>1</LanguageType>
            </Language>
          </LanguageList>
          <DwellingArrangement>
            <Code>1671</Code>
          </DwellingArrangement>
          <Religion>
            <Code>2013</Code>
          </Religion>
          <ReligiousEventList>
            <ReligiousEvent>
              <Type>Baptism</Type>
              <Date>2000-09-01</Date>
            </ReligiousEvent>
            <ReligiousEvent>
              <Type>Christmas</Type>
              <Date>2009-12-24</Date>
            </ReligiousEvent>
          </ReligiousEventList>
          <ReligiousRegion>The Religion Region</ReligiousRegion>
          <PermanentResident>P</PermanentResident>
          <VisaSubClass>101</VisaSubClass>
          <VisaStatisticalCode>05</VisaStatisticalCode>
        </Demographics>
        <AddressList>
          <Address Type ="0123" Role="012B">
            <Street>
              <Line1>Unit1/10</Line1>
              <Line2>Barkley Street</Line2>
            </Street>
            <City>Yarra Glenn</City>
            <StateProvince>VIC</StateProvince>
            <Country>1101</Country>
            <PostalCode>9999</PostalCode>
          </Address>
          <Address Type="0123A" Role="013A">
            <Street>
              <Line1>34 Term Address Street</Line1>
            </Street>
            <City>Home Town</City>
            <StateProvince>WA</StateProvince>
            <Country>1101</Country>
            <PostalCode>9999</PostalCode>
          </Address>
        </AddressList>
        <PhoneNumberList>
          <PhoneNumber Type="0096">
            <Number>03 9637-2289</Number>
            <Extension>72289</Extension>
            <ListedStatus>Y</ListedStatus>
          </PhoneNumber>
          <PhoneNumber Type="0888">
            <Number>0437-765-234</Number>
            <ListedStatus>N</ListedStatus>
          </PhoneNumber>
        </PhoneNumberList>
        <EmailList>
          <Email Type="01">fsmith@yahoo.com</Email>
          <Email Type="02">freddy@gmail.com</Email>
        </EmailList>
      </PersonInfo>
     
      <EmploymentType>4</EmploymentType>
      <SchoolEducationalLevel>3</SchoolEducationalLevel>
        <NonSchoolEducation>6</NonSchoolEducation>
		<Employment>Fletcher</Employment>
 		<Workplace>Bob's Arrows</Workplace>
     </StudentContactPersonal>

`
var test_StudentContactRelationship = `    <StudentContactRelationship StudentContactRelationshipRefId = "7572B261-0947-583A-463D-BB345291B332">
      <StudentPersonalRefId>DEE34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <StudentContactPersonalRefId>6472B261-0947-583A-463D-BB345291B001</StudentContactPersonalRefId>
      <Relationship>
        <Code>01</Code>
      </Relationship>
      <HouseholdList>
        <Household>FamilyA</Household>
        <Household>FamilyB</Household>
      </HouseholdList>
      <ContactFlags>
        <ParentLegalGuardian>Y</ParentLegalGuardian>
        <PickupRights>Y</PickupRights>
        <LivesWith>N</LivesWith>
        <AccessToRecords>U</AccessToRecords>
        <ReceivesAssessmentReport>Y</ReceivesAssessmentReport>
        <EmergencyContact>Y</EmergencyContact>
        <HasCustody>N</HasCustody>
        <DisciplinaryContact>N</DisciplinaryContact>
        <AttendanceContact>N</AttendanceContact>
        <PrimaryCareProvider>U</PrimaryCareProvider>
        <FeesBilling>Y</FeesBilling>
        <FamilyMail>Y</FamilyMail>
        <InterventionOrder>N</InterventionOrder>
      </ContactFlags>
      <MainlySpeaksEnglishAtHome>U</MainlySpeaksEnglishAtHome>
        <ContactSequence>1</ContactSequence>
        <ContactSequenceSource>P</ContactSequenceSource>
        <ContactMethod>AltMailing</ContactMethod>
		<FeePercentage>
		  <Curriculum>90.0</Curriculum>
		  <Other>10.0</Other>
		</FeePercentage>
      </StudentContactRelationship>

`
var test_StudentDailyAttendance = `    <StudentDailyAttendance RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
      <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
      <Date>2002-11-01</Date>
      <SchoolYear>2003</SchoolYear>
      <DayValue>Partial</DayValue>
      <AttendanceCode>
        <Code>200</Code>
        <OtherCodeList>
          <OtherCode Codeset="Local">S</OtherCode>
          <OtherCode Codeset="Text">C</OtherCode>
        </OtherCodeList>
      </AttendanceCode>
      <AttendanceStatus>01</AttendanceStatus>
      <TimeIn>13:30:00</TimeIn>
      <TimeOut>12:05:00</TimeOut>
      <AttendanceNote>Left for Orthodontist appt. and returned to school afterward</AttendanceNote>
    </StudentDailyAttendance>

`
var test_StudentGrade = `    <StudentGrade RefId="359D7510-1AD0-A9D7-A8C3-DAD0A85103A2">
      <StudentPersonalRefId>A137D78A-E00B-C744-EF90-F2871CEB90A2</StudentPersonalRefId>
      <SchoolInfoRefId>B237D78A-E00B-C744-EF90-F2871CEB91A2</SchoolInfoRefId>
      <Grade>
        <Percentage>79</Percentage>
        <Numeric>3.0</Numeric>
        <Letter>C</Letter>
        <Narrative>Johnny is only achieving average performance.</Narrative>
      </Grade>
 	  <TermSpan>0828</TermSpan>
	  <SchoolYear>2021</SchoolYear>
   </StudentGrade>
`
var test_StudentParticipation = `    <StudentParticipation RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      <StudentPersonalRefId>A2E34F59-A742-C1A4-B3D1-1CC002B163A2</StudentPersonalRefId>
      <StudentParticipationAsOfDate>2006-07-13</StudentParticipationAsOfDate>
      <ProgramType>0240</ProgramType>
      <ProgramFundingSources>
        <ProgramFundingSource>
          <Code>1</Code>
        </ProgramFundingSource>
      </ProgramFundingSources>
      
      <ManagingSchool SIF_RefObject="SchoolInfo">D93F4D18-3A42-C1A4-B3D1-1CC002B163A2</ManagingSchool>
      
      
      <ParticipationContact>John Mason</ParticipationContact>
    </StudentParticipation>
`
var test_StudentPeriodAttendanceExample = `    <StudentPeriodAttendance RefId="98157AA0-13BA-8C3D-00AA-012B359D7512">
      <StudentPersonalRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</StudentPersonalRefId>
      <SchoolInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</SchoolInfoRefId>
      <Date>2001-03-05</Date>
      <TimetablePeriod>P4</TimetablePeriod>
      <AttendanceCode>
        <Code>802</Code>
      </AttendanceCode>
      <AttendanceStatus>01</AttendanceStatus>
    </StudentPeriodAttendance>
`
var test_example0783 = `		<StudentPersonal RefId="7C834EA9-EDA1-2090-347F-83297E1C290C">
			<AlertMessages>
				<AlertMessage Type="Legal">Mother is legal guardian</AlertMessage>
			</AlertMessages>
			<MedicalAlertMessages>
				<MedicalAlertMessage Severity="Severe">Student has Peanut Allergy</MedicalAlertMessage>
				<MedicalAlertMessage Severity="Moderate">Student has Diabetes</MedicalAlertMessage>
			</MedicalAlertMessages>
			<LocalId>S1234567</LocalId>
			<StateProvinceId>ABC1234</StateProvinceId>
			<NationalUniqueStudentIdentifier>12345678</NationalUniqueStudentIdentifier>
			<ElectronicIdList>
				<ElectronicId Type="03">ZZZZZZ21</ElectronicId>
				<ElectronicId Type="03">ZZZZZZ22</ElectronicId>
			</ElectronicIdList>
			<OtherIdList>
					 <OtherId Type="PreviousNAPPlatformStudentId">888rdgf</OtherId>
					 <OtherId Type="DiocesanStudentId">1234</OtherId>
    		</OtherIdList>
			<PersonInfo>
				<Name Type="LGL">
					<FamilyName>Smith</FamilyName>
					<GivenName>Fred</GivenName>
					<FullName>Fred Smith</FullName>
				</Name>
				<OtherNames>
					<Name Type="AKA">
						<FamilyName>Anderson</FamilyName>
						<GivenName>Samuel</GivenName>
						<FullName>Samuel Anderson</FullName>
					</Name>
					<Name Type="PRF">
						<FamilyName>Rowinski</FamilyName>
						<GivenName>Sam</GivenName>
						<FullName>Sam Rowinski </FullName>
					</Name>
				</OtherNames>
				<Demographics>
					<IndigenousStatus>3</IndigenousStatus>
					<Sex>1</Sex>
					<BirthDate>1990-09-26</BirthDate>
					<BirthDateVerification>1004</BirthDateVerification>
					<PlaceOfBirth>Clayton</PlaceOfBirth>
					<StateOfBirth>VIC</StateOfBirth>
					<CountryOfBirth>1101</CountryOfBirth>
					<CountriesOfCitizenship>
						<CountryOfCitizenship>8104</CountryOfCitizenship>
						<CountryOfCitizenship>1101</CountryOfCitizenship>
					</CountriesOfCitizenship>
					<CountriesOfResidency>
						<CountryOfResidency>8104</CountryOfResidency>
						<CountryOfResidency>1101</CountryOfResidency>
					</CountriesOfResidency>
					<CountryArrivalDate>1990-09-26</CountryArrivalDate>
					<AustralianCitizenshipStatus>1</AustralianCitizenshipStatus>
					<EnglishProficiency>
						<Code>1</Code>
					</EnglishProficiency>
					<LanguageList>
						<Language>
							<Code>0001</Code>
							<LanguageType>1</LanguageType>
						</Language>
					</LanguageList>
					<DwellingArrangement>
						<Code>1671</Code>
					</DwellingArrangement>
					<Religion>
						<Code>2013</Code>
					</Religion>
					<ReligiousEventList>
						<ReligiousEvent>
							<Type>Baptism</Type>
							<Date>2000-09-01</Date>
						</ReligiousEvent>
						<ReligiousEvent>
							<Type>Christmas</Type>
							<Date>2009-12-24</Date>
						</ReligiousEvent>
					</ReligiousEventList>
					<ReligiousRegion>The Religion Region</ReligiousRegion>
					<PermanentResident>P</PermanentResident>
					<VisaSubClass>101</VisaSubClass>
					<VisaStatisticalCode>05</VisaStatisticalCode>
					<VisaSubClassList>
						  <VisaSubClass>
							<Code>401</Code>
							<VisaExpiryDate>2019-07-25</VisaExpiryDate>
							<ATEExpiryDate>2016-12-26</ATEExpiryDate>
							<ATEStartDate>2017-12-26</ATEStartDate>
							<VisaStatisticalCode>05</VisaStatisticalCode>
						  </VisaSubClass>
					</VisaSubClassList>
					<Passport>
						<Number>9</Number>
						<ExpiryDate>2030-12-11</ExpiryDate>
						<Country>4111</Country>
					</Passport>
				</Demographics>
				<AddressList>
					<Address Type="0123" Role="2382">
						<Street>
							<Line1>Unit1/10</Line1>
							<Line2>Barkley Street</Line2>
						</Street>
						<City>Yarra Glenn</City>
						<StateProvince>VIC</StateProvince>
						<Country>1101</Country>
						<PostalCode>9999</PostalCode>
					</Address>
					<Address Type="0123A" Role="013A">
						<Street>
							<Line1>34 Term Address Street</Line1>
						</Street>
						<City>Home Town</City>
						<StateProvince>WA</StateProvince>
						<Country>1101</Country>
						<PostalCode>9999</PostalCode>
					</Address>
				</AddressList>
				<PhoneNumberList>
					<PhoneNumber Type="0096">
						<Number>03 9637-2289</Number>
						<Extension>72289</Extension>
						<ListedStatus>Y</ListedStatus>
					</PhoneNumber>
					<PhoneNumber Type="0888">
						<Number>0437-765-234</Number>
						<ListedStatus>N</ListedStatus>
					</PhoneNumber>
				</PhoneNumberList>
				<EmailList>
					<Email Type="01">fsmith@yahoo.com</Email>
					<Email Type="02">freddy@gmail.com</Email>
				</EmailList>
			</PersonInfo>
			<ProjectedGraduationYear>2014</ProjectedGraduationYear>
			<OnTimeGraduationYear>2012</OnTimeGraduationYear>
			<MostRecent>
				     <SchoolLocalId>S1234567</SchoolLocalId>
				     <HomeroomLocalId>hr12345</HomeroomLocalId>
					 <YearLevel>
					   <Code>P</Code>
					 </YearLevel>
					<FTE>0.5</FTE>
					 <Parent1Language>1201</Parent1Language>
					 <Parent2Language>1201</Parent2Language>
					 <LocalCampusId>D</LocalCampusId>
					 <SchoolACARAId>VIC687</SchoolACARAId>
					 <Homegroup>7A</Homegroup>
					 <ClassCode>English 7D</ClassCode>
					 <MembershipType>02</MembershipType>
					 <FFPOS>2</FFPOS>
					 <ReportingSchoolId>VIC670</ReportingSchoolId>
					 <OtherEnrollmentSchoolACARAId>VIC6273</OtherEnrollmentSchoolACARAId>
			</MostRecent>
			<AcceptableUsePolicy>Y</AcceptableUsePolicy>
			<EconomicDisadvantage>N</EconomicDisadvantage>
			<ESL>Y</ESL>
			<ESLDateAssessed>2016-04-23</ESLDateAssessed>
			<YoungCarersRole>N</YoungCarersRole>
			<Disability>N</Disability>
			<IntegrationAide>N</IntegrationAide>
			<EducationSupport>N</EducationSupport>
			<HomeSchooledStudent>N</HomeSchooledStudent>
			<IndependentStudent>N</IndependentStudent>
			<Sensitive>N</Sensitive>
			<LocalCodeList>
				<LocalCode>
					<LocalisedCode>0921</LocalisedCode>
					<Description>Serbia Montenegro</Description>
					<Element>StudentPersonal/PersonInfo/Demographics/CountryOfBirth</Element>
				</LocalCode>
			    <LocalCode>
					<LocalisedCode>0921</LocalisedCode>
					<Description>Serbia Montenegro</Description>
					<Element>StudentPersonal/PersonInfo/Demographics/CountriesOfCitizenship/CountryOfCitizenship</Element>
					<ListIndex>1</ListIndex> 
				</LocalCode>
			</LocalCodeList>
		</StudentPersonal>
`
var test_StudentSchoolEnrollment = `    <StudentSchoolEnrollment RefId="A8C3D3E3-4B35-9D75-101D-00AA001A1652">
      <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <SchoolInfoRefId>D3E34B35-9D75-101A-8C3D-00AA001A1651</SchoolInfoRefId>
      <MembershipType>01</MembershipType>
      <TimeFrame>C</TimeFrame>
      <SchoolYear>2004</SchoolYear>
      <IntendedEntryDate>2004-01-28</IntendedEntryDate>
      <EntryDate>2004-01-29</EntryDate>
      <EntryType>
        <Code>1838</Code>
      </EntryType>
      <YearLevel>
        <Code>10</Code>
      </YearLevel>
      <Homeroom SIF_RefObject="RoomInfo">D7510D3E-34B3-591A-8C3D-00AA001A1651</Homeroom>
      <Advisor SIF_RefObject="StaffPersonal">B359D3E3-4D75-101A-8C3D-00AA001A1652</Advisor>
      <Counselor SIF_RefObject="StaffPersonal">B459D3E3-4D75-101A-8C3D-00AA001A1653</Counselor>
      <ExitType>
        <Code>1909</Code>
        <OtherCodeList>
            <OtherCode Codeset="Local">Victoria</OtherCode>
        </OtherCodeList>
	  </ExitType>
      <FTE>1.00</FTE>
      <FTPTStatus>01</FTPTStatus>
      <PublishingPermissionList>
		  <PublishingPermission>
			  <PermissionCategory>OKMediaRelease</PermissionCategory>
			  <PermissionValue>N</PermissionValue>
		  </PublishingPermission>
		  <PublishingPermission>
			<PermissionCategory>OKPrintedMaterial</PermissionCategory>
			<PermissionValue>Y</PermissionValue>
		  </PublishingPermission>
	  </PublishingPermissionList>
      <CensusAge>16</CensusAge>
      <InternationalStudent>Y</InternationalStudent>
      <TravelDetails>
		  <ToSchool>
			  <TravelMode>A</TravelMode> 
			  <TravelDetails>Usually walks to school unless otherwise advised.</TravelDetails>
			  <TravelAccompaniment>I</TravelAccompaniment>
   		  </ToSchool>	  
		  <FromSchool>
			  <TravelMode>A</TravelMode>
			  <TravelDetails>Usually walks from school unless otherwise advised.</TravelDetails>
			  <TravelAccompaniment>I</TravelAccompaniment>
		  </FromSchool>
      </TravelDetails>
    </StudentSchoolEnrollment>
`
var test_StudentScoreJudgementAgainstStandardExample1ObjectsPublishedpreviously = `    <StudentScoreJudgementAgainstStandard RefId="5810E283-E928-459C-BBA7-6EFE1963F784">
      <SchoolYear>2018</SchoolYear>
      <TermInfoRefId>1F971F3D-414E-4146-9325-66638290E6C1</TermInfoRefId>
      <StudentPersonalRefId>8F0934CC-2F04-48F8-BBD7-88AA1ADE691B</StudentPersonalRefId>
      <YearLevel>
		  <Code>10</Code>
      </YearLevel>
      <TeachingGroupRefId>E5EDAE63-A04D-47E4-9D02-24A32956B074</TeachingGroupRefId>
      <StaffPersonalRefId>9D0934CC-2F04-48F8-BBD7-88AA1ADE691B</StaffPersonalRefId>
      <CurriculumCode>AUSVELS</CurriculumCode>
      <CurriculumNodeCode>MATNUM</CurriculumNodeCode>
      <Score>10</Score>
      <SpecialCircumstanceLocalCode>SCMP</SpecialCircumstanceLocalCode>
      <ManagedPathwayLocalCode>MPOMP</ManagedPathwayLocalCode>
      <SchoolInfoRefId>E6EDAE63-A04D-47E4-9D02-24A32956B075</SchoolInfoRefId>
    </StudentScoreJudgementAgainstStandard>
`
var test_StudentScoreJudgementAgainstStandardExample2NopreviousSIFObjectsPublished = `    <StudentScoreJudgementAgainstStandard RefId="5810E283-E928-459C-BBA7-6EFE1963F784">
      <SchoolYear>2018</SchoolYear>
      <LocalTermCode>Mid</LocalTermCode>
      <StudentStateProvinceId>987654321</StudentStateProvinceId>
      <YearLevel>
		  <Code>10</Code>
      </YearLevel>
      <ClassLocalId>MAT10A</ClassLocalId>
      <StaffLocalId>C1234567</StaffLocalId>
      <CurriculumCode>AUSVELS</CurriculumCode>
      <CurriculumNodeCode>MATNUM</CurriculumNodeCode>
      <Score>10</Score>
      <SpecialCircumstanceLocalCode>SCMP</SpecialCircumstanceLocalCode>
      <ManagedPathwayLocalCode>MPOMP</ManagedPathwayLocalCode>
      <SchoolLocalId>027789</SchoolLocalId>
      <SchoolCommonwealthId>AGEID-1032</SchoolCommonwealthId>
    </StudentScoreJudgementAgainstStandard>
`
var test_StudentSectionEnrollment = `    <StudentSectionEnrollment RefId="983AC165-9879-3002-C3D0-0AA00456789D" >
	  <StudentPersonalRefId>CAE29316-5987-101A-8C3D-00AA00456789</StudentPersonalRefId>
      <SectionInfoRefId>9076AB23-E386-112B-7EA2-256100BB3312</SectionInfoRefId>
	  <SchoolYear>2013</SchoolYear>
      <EntryDate>2013-02-02</EntryDate>
      <ExitDate>2013-11-15</ExitDate>
      
    </StudentSectionEnrollment>
`
var test_TeachingGroup = `    <TeachingGroup RefId="64A309DA-063A-2E35-B359-D75101A8C3D1">
      <SchoolYear>2008</SchoolYear>
      <LocalId>20087ASPN</LocalId>
      <ShortName>7A SPN</ShortName>
      <LongName>Year 7A Maths - Space and Numbers</LongName>
      <Set>4</Set>
      <Block>6</Block>
      <CurriculumLevel>VELS Level 5</CurriculumLevel>
      <StudentList>
        <TeachingGroupStudent>
          <StudentPersonalRefId>9897466F-200E-4BC1-B9AE-D1507DA15CEF</StudentPersonalRefId>
          <StudentLocalId>SMI001</StudentLocalId>
          <Name Type="LGL">
            <FamilyName>Smith</FamilyName>
            <GivenName>Peter</GivenName>
          </Name>
        </TeachingGroupStudent>
        <TeachingGroupStudent>
          <StudentPersonalRefId>7C834EA9-EDA1-2090-347F-83297E1C290D</StudentPersonalRefId>
          <StudentLocalId>SMI002</StudentLocalId>
          <Name Type="LGL">
            <FamilyName>Smith</FamilyName>
            <GivenName>Jennifer</GivenName>
          </Name>
        </TeachingGroupStudent>
        <TeachingGroupStudent>
          <StudentPersonalRefId>7C834EA9-EDA1-2090-347F-83297E1C290E</StudentPersonalRefId>
          <StudentLocalId>SMI003</StudentLocalId>
          <Name Type="LGL">
            <FamilyName>Smith</FamilyName>
            <GivenName>Terence</GivenName>
          </Name>
        </TeachingGroupStudent>
      </StudentList>
      <TeacherList>
        <TeachingGroupTeacher>
          <StaffPersonalRefId>A8C3A2E3-5B35-9D75-101D-00AA001A0000</StaffPersonalRefId>
          <StaffLocalId>SMI1</StaffLocalId>
          <Name Type="LGL">
            <FamilyName>Smith</FamilyName>
            <GivenName>Thomas</GivenName>
          </Name>
          <Association>Class Teacher</Association>
        </TeachingGroupTeacher>
        <TeachingGroupTeacher>
          <StaffPersonalRefId>A8CCCCE3-5B35-9D75-101D-00AA001A0000</StaffPersonalRefId>
          <StaffLocalId>LONG2</StaffLocalId>
          <Name Type="LGL">
            <FamilyName>Long</FamilyName>
            <GivenName>Tamara</GivenName>
           
          </Name>
          <Association>Integration Aide</Association>
        </TeachingGroupTeacher>
      </TeacherList>
      <TeachingGroupPeriodList>
         <TeachingGroupPeriod>
			 <DayId>M</DayId>
			 <PeriodId>2</PeriodId>
         </TeachingGroupPeriod>
         <TeachingGroupPeriod>
			 <DayId>F</DayId>
			 <PeriodId>6</PeriodId>
         </TeachingGroupPeriod>
      </TeachingGroupPeriodList>
    </TeachingGroup>

    
`
var test_TermInfo = `    <TermInfo RefId="7E59D751-01A8-0A70-0163-75DE097A0726">
      <SchoolInfoRefId>A2E35B35-9D75-101A-8C3D-00AA001A0000</SchoolInfoRefId>
      <SchoolYear>2012</SchoolYear>
      <StartDate>2012-01-05</StartDate>
      <EndDate>2012-04-01</EndDate>
      <Description>First Term 2012</Description>
      <RelativeDuration>0.2500</RelativeDuration>
      <TermCode>Sp2004</TermCode>
      <Track>PreK</Track>
      <TermSpan>0833</TermSpan>
    </TermInfo>
`
var test_TimeTable = `    <TimeTable RefId="64A309DA-063A-2E35-B359-D75101A8C3D0">
      <SchoolInfoRefId>F2256EE2-B67F-47D6-AB47-94D4DEE0D0AD</SchoolInfoRefId>
      <SchoolYear>2008</SchoolYear>
      <LocalId>2008S1</LocalId>
      <Title>2008 Semester 1</Title>
      <DaysPerCycle>5</DaysPerCycle>
      <PeriodsPerDay>6</PeriodsPerDay>
      <TeachingPeriodsPerDay>5</TeachingPeriodsPerDay>
      <SchoolLocalId>01011234</SchoolLocalId>
      <SchoolName>Newest Secondary College</SchoolName>
      <TimeTableCreationDate>2008-02-01</TimeTableCreationDate>
      <StartDate>2008-01-30</StartDate>
      <EndDate>2008-06-20</EndDate>
      <TimeTableDayList>
        <TimeTableDay>
          <DayId>1</DayId>
          <DayTitle>Monday</DayTitle>
          <TimeTablePeriodList>
            <TimeTablePeriod>
              <PeriodId>1</PeriodId>
              <PeriodTitle>P1</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>2</PeriodId>
              <PeriodTitle>P2</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>3</PeriodId>
              <PeriodTitle>P3</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>4</PeriodId>
              <PeriodTitle>P4</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>5</PeriodId>
              <PeriodTitle>P5</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>6</PeriodId>
              <PeriodTitle>P6</PeriodTitle>
            </TimeTablePeriod>
          </TimeTablePeriodList>
        </TimeTableDay>
        <TimeTableDay>
          <DayId>2</DayId>
          <DayTitle>Tuesday</DayTitle>
          <TimeTablePeriodList>
            <TimeTablePeriod>
              <PeriodId>1</PeriodId>
              <PeriodTitle>P1</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>2</PeriodId>
              <PeriodTitle>P2</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>3</PeriodId>
              <PeriodTitle>P3</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>4</PeriodId>
              <PeriodTitle>P4</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>5</PeriodId>
              <PeriodTitle>P5</PeriodTitle>
            </TimeTablePeriod>
            <TimeTablePeriod>
              <PeriodId>6</PeriodId>
              <PeriodTitle>P6</PeriodTitle>
            </TimeTablePeriod>
          </TimeTablePeriodList>
        </TimeTableDay>
        
      </TimeTableDayList>
    </TimeTable>

`
var test_TimeTableCell = `    <TimeTableCell RefId = "64A309DA-063A-2E35-B359-D75101A8C3D1">
   <TimeTableRefId>64A309DA-063A-2E35-B359-D75101A8C3D0</TimeTableRefId>
      <TimeTableSubjectRefId>22860091-7192-45B4-AB0C-F5B9DC19DE5C</TimeTableSubjectRefId>
      <TeachingGroupRefId>64A309DA-063A-2E35-B359-D75101A8C3D1</TeachingGroupRefId>
      <RoomInfoRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</RoomInfoRefId>
      <StaffPersonalRefId>D3E34F41-9D75-101A-8C3D-00AA001A1652</StaffPersonalRefId>
      <TimeTableLocalId>2008S1</TimeTableLocalId>
      <SubjectLocalId>07AR</SubjectLocalId>
        <TeachingGroupLocalId>20087ASPN 2008S1</TeachingGroupLocalId>
        <RoomNumber>101</RoomNumber>
        <StaffLocalId>946379881</StaffLocalId>
        <DayId>1</DayId>
        <PeriodId>1</PeriodId>
        <CellType>T</CellType>
        <SchoolInfoRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</SchoolInfoRefId>
        <SchoolLocalId>01011234</SchoolLocalId>
            
         <TeacherList>
			 <TeacherCover>
			   <StaffPersonalRefId>98157AA0-13BA-8C3D-00AA-012B359D7512</StaffPersonalRefId>
			   <StaffLocalId>US8923</StaffLocalId>
			   <StartTime>12:05:00</StartTime>
			   <FinishTime>13:00:00</FinishTime>
			   <Credit>In-Lieu</Credit>
			   <Supervision>MinimalSupervision</Supervision>
			   <Weighting>0.5</Weighting>
			 </TeacherCover>
			 <TeacherCover>
			   <StaffPersonalRefId>A75A0010-1A8C-301D-02E3-A05B359D0A00</StaffPersonalRefId>
			   <StaffLocalId>ZY3213</StaffLocalId>
			   <StartTime>13:00:00</StartTime>
			   <FinishTime>13:30:00</FinishTime>
			   <Credit>Extra</Credit>
			   <Supervision>MinimalSupervision</Supervision>
			   <Weighting>0.5</Weighting>
			 </TeacherCover>
		</TeacherList>
		<RoomList>
			 <RoomInfoRefId>11737EA4-301C-ADCA-75C8-7214A7C46BDB</RoomInfoRefId>
		</RoomList>            
    </TimeTableCell>

`
var test_TimeTableContainer = ``
var test_TimeTableSubject = `    <TimeTableSubject RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      <SubjectLocalId>07AR</SubjectLocalId>
      <AcademicYear>
        <Code>7</Code>
      </AcademicYear>
      <SubjectShortName>7 ART</SubjectShortName>
      <SubjectLongName>Year 7 Art</SubjectLongName>
      <SubjectType>E</SubjectType>
      <SchoolYear>2009</SchoolYear>
    </TimeTableSubject>

`
var test_WellbeingAlert = `    <WellbeingAlert RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
      <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
      <Date>2017-11-01</Date>
      <WellbeingAlertStartDate>2017-11-01</WellbeingAlertStartDate>
      <WellbeingAlertCategory>M</WellbeingAlertCategory>
      <WellbeingAlertDescription>This Student is allergic to peanuts</WellbeingAlertDescription>
    </WellbeingAlert>

`
var test_WellbeingAppeal = `    <WellbeingAppeal RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
      <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
      <WellbeingResponseRefId>BC385746-359D-7510-1A8C-36432A901A36</WellbeingResponseRefId>
      <AppealStatusCode>SU</AppealStatusCode>
      <Date>2017-11-01</Date>
      <AppealNotes>This Appeal was successful.</AppealNotes>
    </WellbeingAppeal>

`
var test_WellbeingCharacteristic = `    <WellbeingCharacteristic RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
      <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
      <WellbeingCharacteristicStartDate>2017-11-01</WellbeingCharacteristicStartDate>
      <WellbeingCharacteristicEndDate>2017-11-10</WellbeingCharacteristicEndDate>
      <WellbeingCharacteristicNotes>This Student has a special need.</WellbeingCharacteristicNotes>
      <WellbeingCharacteristicCategory>S</WellbeingCharacteristicCategory>
	  <PreferredHospital>St. Aloysius'</PreferredHospital>
    </WellbeingCharacteristic>

`
var test_WellbeingEvent = `    <WellbeingEvent RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
      <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
      <WellbeingEventCategoryClass>P</WellbeingEventCategoryClass>
      <WellbeingEventDate>2017-11-01</WellbeingEventDate>
      <WellbeingEventTimePeriod>AM</WellbeingEventTimePeriod>
    </WellbeingEvent>

`
var test_WellbeingEvent2 = `	  <WellbeingEvent RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
	  <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
	  <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
	  <EventId>Yard-10</EventId>
	  <WellbeingEventNotes>Incident relates to School Yard accident.</WellbeingEventNotes>
	  <WellbeingEventCategoryClass>D</WellbeingEventCategoryClass>
	  <WellbeingEventCategoryList>
		  <WellbeingEventCategory>
			  <EventCategory>M5</EventCategory>
			  <WellbeingEventSubCategoryList>
				  <WellbeingEventSubCategory>M5-a</WellbeingEventSubCategory>
			  </WellbeingEventSubCategoryList>
		  </WellbeingEventCategory>
	  </WellbeingEventCategoryList>
	  <ReportingStaffRefId>3FFB63B4-CFEF-4820-8501-E7D1E54555CB</ReportingStaffRefId>
	  <WellbeingEventDate>2017-11-01</WellbeingEventDate>
	  <WellbeingEventTimePeriod>AM</WellbeingEventTimePeriod>
	</WellbeingEvent>
`
var test_WellbeingEvent3 = `		<WellbeingEvent RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
			<GroupIndicator>true</GroupIndicator>
			<SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
			<EventId>Yard-10</EventId>
			<WellbeingEventNotes>Incident relates to School Yard accident.</WellbeingEventNotes>
			<WellbeingEventCategoryClass>D</WellbeingEventCategoryClass>
			<WellbeingEventCategoryList>
				<WellbeingEventCategory>
					<EventCategory>M5</EventCategory>
					<WellbeingEventSubCategoryList>
						<WellbeingEventSubCategory>M5-a</WellbeingEventSubCategory>
					</WellbeingEventSubCategoryList>
				</WellbeingEventCategory>
			</WellbeingEventCategoryList>
			<ReportingStaffRefId>3FFB63B4-CFEF-4820-8501-E7D1E54555CB</ReportingStaffRefId>
			<WellbeingEventDate>2017-11-01</WellbeingEventDate>
			<WellbeingEventTimePeriod>AM</WellbeingEventTimePeriod>
		</WellbeingEvent>
	
`
var test_WellbeingPersonLink = `		<WellbeingPersonLink RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
			<WellbeingEventRefId>D3E34B35-9D75-101A-8C3D-00AA001A1688</WellbeingEventRefId>
			<WellbeingResponseRefId>E6E34B35-9D75-101A-8C3D-00AA001A1700</WellbeingResponseRefId>
			<GroupId>Inv0009</GroupId>
			<PersonRefId SIF_RefObject="StudentPersonal">D3E34B35-9D75-101A-8C3D-00AA001A1652</PersonRefId>
			<!-- Should you have the school here???
			<SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>-->
			<HowInvolved>Victim</HowInvolved>
			<FollowUpActionList>
			 <FollowUpAction>
				 <FollowUpDetails>Parent's notified</FollowUpDetails>
			 </FollowUpAction>
			</FollowUpActionList>
		</WellbeingPersonLink>
`
var test_WellbeingPersonLink2 = `		<WellbeingPersonLink RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
			<WellbeingEventRefId>D3E34B35-9D75-101A-8C3D-00AA001A1688</WellbeingEventRefId>
			<WellbeingResponseRefId>E6E34B35-9D75-101A-8C3D-00AA001A1700</WellbeingResponseRefId>
			<GroupId>Inv0009</GroupId>
			<PersonRefId SIF_RefObject="StudentPersonal">D3E34B35-9D75-101A-8C3D-00AA001A1652</PersonRefId>
			<!-- Should you have the school here???
			<SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>-->
			<HowInvolved>Instigator</HowInvolved>
			<FollowUpActionList>
			 <FollowUpAction>
				 <FollowUpDetails>Parent's notified</FollowUpDetails>
			 </FollowUpAction>
			</FollowUpActionList>
		</WellbeingPersonLink>
`
var test_WellbeingPersonLink3 = `		<WellbeingPersonLink RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
			<WellbeingEventRefId>D3E34B35-9D75-101A-8C3D-00AA001A1688</WellbeingEventRefId>
			<WellbeingResponseRefId>E6E34B35-9D75-101A-8C3D-00AA001A1700</WellbeingResponseRefId>
			<GroupId>Inv0009</GroupId>
			<PersonRefId SIF_RefObject="StaffPersonal">D3E34B35-9D75-101A-8C3D-00AA001A1652</PersonRefId>
			<HowInvolved>Reporter</HowInvolved>
			<FollowUpActionList>
			 <FollowUpAction>
				 <FollowUpDetails>Event Logged</FollowUpDetails>
			 </FollowUpAction>
			</FollowUpActionList>
		</WellbeingPersonLink>
`
var test_WellbeingPersonLink4 = `		<WellbeingPersonLink RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
			<WellbeingEventRefId>D3E34B35-9D75-101A-8C3D-00AA001A1688</WellbeingEventRefId>
			<WellbeingResponseRefId>E6E34B35-9D75-101A-8C3D-00AA001A1700</WellbeingResponseRefId>
			<GroupId>Inv0009</GroupId>
			<OtherPersonId>Not known</OtherPersonId>
			<OtherPersonContactDetails>Ms Smith, Visitor, 0435 345 678</OtherPersonContactDetails>
			<PersonRole>Reported Incident to Office</PersonRole>
		</WellbeingPersonLink>
`
var test_WellbeingResponse = `    <WellbeingResponse RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
      <StudentPersonalRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</StudentPersonalRefId>
      <SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId>
      <Date>2017-11-01</Date>
      <WellbeingResponseStartDate>2017-11-01</WellbeingResponseStartDate>
      <WellbeingResponseEndDate>2017-11-10</WellbeingResponseEndDate>
      <WellbeingResponseCategory>S</WellbeingResponseCategory>
      <WellbeingResponseNotes>This Student is to be withdrawn.</WellbeingResponseNotes>
    </WellbeingResponse>

`
var test_AddressCollection = `      
    <AddressCollection RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
        <AddressCollectionYear>2019</AddressCollectionYear> 
        <RoundCode>SES1</RoundCode>
        <SoftwareVendorInfo>  
            <SoftwareProduct>Civica</SoftwareProduct>
            <SoftwareVersion>Websys 1.0</SoftwareVersion>
        </SoftwareVendorInfo>
			<AddressCollectionReportingList> 
                    <AddressCollectionReporting>
                          <CommonwealthId>012345</CommonwealthId> 
                          <EntityName>XXX Secondary College</EntityName>
                          <EntityContact>
                                 <Name Type="LGL">
                                   <FamilyName>Citizen</FamilyName> 
                                   <GivenName>John</GivenName> 
                                 </Name>
                                 <PositionTitle>Auditor</PositionTitle> 
                                 <Email Type="01">jcitizen@school.vic.edu.au</Email> 
                                 <PhoneNumber Type="0096"><Number>03 1234-2678</Number></PhoneNumber> 
                          </EntityContact>
                          <AGContextualQuestionList>
							  <AGContextualQuestion>
								  <AGContextCode>AC001</AGContextCode>
								  <AGAnswer>Yes</AGAnswer>
							  </AGContextualQuestion>
							  <AGContextualQuestion>
   								  <AGContextCode>AC002</AGContextCode>
								  <AGAnswer>Primary</AGAnswer>
							  </AGContextualQuestion>
                          </AGContextualQuestionList>
                          <AddressCollectionStudentList>
                              <AddressCollectionStudent>
                                       <LocalId>12345</LocalId>
                                       <EducationLevel>Primary</EducationLevel>      
                                       <BoardingStatus>D</BoardingStatus>  
                                       <ReportingParent2>Y</ReportingParent2>                                         
                                       <StudentAddress Type="0123" Role="012B">
											<Street>
												<Line1>The House</Line1>
												<Line2>2921 Warburton Hwy</Line2>
											</Street>
											<City>Chicago</City>
											<StateProvince>WA</StateProvince>
											<PostalCode>60611</PostalCode>
										</StudentAddress>
										<Parent1>
										    <ParentName Type="LGL">
											   <FamilyName>Miller</FamilyName> 
											   <GivenName>James</GivenName> 
											   <PreferredGivenName>Jim</PreferredGivenName> 
											</ParentName>
											<AddressSameAsStudent>Y</AddressSameAsStudent>
											<ParentAddress Type="0123" Role="012B">
												<Street>
													<Line1>The House</Line1>
													<Line2>2921 Warburton Hwy</Line2>
												</Street>
												<City>Chicago</City>
												<StateProvince>WA</StateProvince>
												<PostalCode>60611</PostalCode>
											</ParentAddress>
										</Parent1>	
										<Parent2>
										    <ParentName Type="LGL">
											   <FamilyName>Miller</FamilyName> 
											   <GivenName>Samantha</GivenName> 
											   <PreferredGivenName>Sam</PreferredGivenName> 
											</ParentName>
											<AddressSameAsStudent>Y</AddressSameAsStudent>
											<ParentAddress Type="0123" Role="012B">
												<Street>
													<Line1>The House</Line1>
													<Line2>2921 Warburton Hwy</Line2>
												</Street>
												<City>Chicago</City>
												<StateProvince>WA</StateProvince>
												<PostalCode>60611</PostalCode>
											</ParentAddress>
										</Parent2>
                              	</AddressCollectionStudent>
                          	</AddressCollectionStudentList >
              		</AddressCollectionReporting>
         		</AddressCollectionReportingList>
       </AddressCollection>

`
var test_CensusCollection = `      
    <CensusCollection RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
		<CensusYear>2018</CensusYear> 
		<RoundCode>Census1</RoundCode>
		<SoftwareVendorInfo>  
			<SoftwareProduct>Civica</SoftwareProduct>
			<SoftwareVersion>Websys 1.0</SoftwareVersion>
		</SoftwareVendorInfo>
		<CensusReportingList> 
			<CensusReporting>
				<EntityLevel>School</EntityLevel>
 				<CommonwealthId>012345</CommonwealthId> 
				<EntityName>XXX Secondary College</EntityName>
				<EntityContact>
					<Name Type="LGL">
					  <FamilyName>Citizen</FamilyName> 
					  <GivenName>James</GivenName> 
					</Name>
					<PositionTitle>Auditor</PositionTitle> 
					<Email Type="01">jcitizen@school.vic.edu.au</Email> 
					<PhoneNumber Type="0096"><Number>03 1234-2678</Number></PhoneNumber> 
				</EntityContact>
				<CensusStaffList>
				    <CensusStaff>
						<StaffCohortId>1</StaffCohortId>
						<StaffActivity><Code>1100</Code></StaffActivity>  
						<CohortGender>M</CohortGender>
						<CohortIndigenousType>N</CohortIndigenousType>
						<PrimaryFTE>0.5</PrimaryFTE>									
						<SecondaryFTE>0.5</SecondaryFTE>			
						<JobFTE>1.0</JobFTE>
						<Headcount>5</Headcount>
				    </CensusStaff>
					<CensusStaff>
						<StaffCohortId>2</StaffCohortId>
						<StaffActivity><Code>1100</Code></StaffActivity>
						<CohortGender>F</CohortGender>
						<CohortIndigenousType>N</CohortIndigenousType>
						<PrimaryFTE>1.0</PrimaryFTE>									
						<SecondaryFTE>0</SecondaryFTE>			
						<JobFTE>1.0</JobFTE>
						<Headcount>6</Headcount>
				    </CensusStaff>	
				    <CensusStaff>
						<StaffCohortId>3</StaffCohortId>
						<StaffActivity><Code>1100</Code></StaffActivity>
						<CohortGender>M</CohortGender>
						<CohortIndigenousType>I</CohortIndigenousType>
  						<PrimaryFTE>0</PrimaryFTE>									
						<SecondaryFTE>1.0</SecondaryFTE>			
						<JobFTE>1.0</JobFTE>
						<Headcount>6</Headcount>
				    </CensusStaff>				
					<CensusStaff>
						<StaffCohortId>4</StaffCohortId>
						<StaffActivity><Code>1103</Code></StaffActivity>
						<CohortGender>M</CohortGender> 
						<CohortIndigenousType>I</CohortIndigenousType>	 
						<PrimaryFTE>0.0</PrimaryFTE>									
						<SecondaryFTE>1.0</SecondaryFTE>			
						<JobFTE>1.0</JobFTE>
						<Headcount>1</Headcount>
				    </CensusStaff>	
				</CensusStaffList>
				<CensusStudentList>
					<CensusStudent>
						<StudentCohortId>1</StudentCohortId>
						<YearLevel><Code>7</Code></YearLevel>	
						<CensusAge>13</CensusAge>	
						<CohortGender>M</CohortGender>					 
						<CohortIndigenousType>I</CohortIndigenousType>	
						<EducationMode>D</EducationMode>	 
						<StudentOnVisa>N</StudentOnVisa>
						<OverseasStudent>Y</OverseasStudent>
						<DisabilityLevelOfAdjustment>None</DisabilityLevelOfAdjustment>  
						<DisabilityCategory>None</DisabilityCategory>  
						<FTE>1.0</FTE>								 
						<Headcount>1</Headcount>
					</CensusStudent>
					<CensusStudent>
						<StudentCohortId>2</StudentCohortId>
						<YearLevel><Code>7</Code></YearLevel>	
						<CensusAge>13</CensusAge>	
						<CohortGender>F</CohortGender>					 
						<CohortIndigenousType>I</CohortIndigenousType>	
						<EducationMode>D</EducationMode>	 
						<StudentOnVisa>N</StudentOnVisa>  
						<OverseasStudent>Y</OverseasStudent>	 
						<DisabilityLevelOfAdjustment>None</DisabilityLevelOfAdjustment>  
						<DisabilityCategory>None</DisabilityCategory>  
						<FTE>1.0</FTE>								 
						<Headcount>1</Headcount>
					</CensusStudent>	
					<CensusStudent>
						<StudentCohortId>3</StudentCohortId>
						<YearLevel><Code>7</Code></YearLevel>
						<CensusAge>13</CensusAge>
						<CohortGender>M</CohortGender>					 
						<CohortIndigenousType>N</CohortIndigenousType>	
						<EducationMode>D</EducationMode>	 
						<StudentOnVisa>N</StudentOnVisa>  
						<OverseasStudent>N</OverseasStudent>	 
						<DisabilityLevelOfAdjustment>None</DisabilityLevelOfAdjustment>  
						<DisabilityCategory>None</DisabilityCategory>  
						<FTE>1.0</FTE>								
						<Headcount>120</Headcount>
					</CensusStudent>
					<CensusStudent>
						<StudentCohortId>4</StudentCohortId>
						<YearLevel><Code>7</Code></YearLevel>		 
						<CensusAge>13</CensusAge>	
						<CohortGender>M</CohortGender>					 
						<CohortIndigenousType>I</CohortIndigenousType>
						<EducationMode>D</EducationMode>	 
						<StudentOnVisa>N</StudentOnVisa> 
						<OverseasStudent>N</OverseasStudent>	 
						<DisabilityLevelOfAdjustment>Substantial</DisabilityLevelOfAdjustment>  
						<DisabilityCategory>Sensory</DisabilityCategory>  
						<FTE>1.0</FTE>								 
						<Headcount>1</Headcount>
					</CensusStudent>
				</CensusStudentList>
		  </CensusReporting>
	  </CensusReportingList>
	</CensusCollection>
   
`
var test_ChargedLocation = `    <ChargedLocationInfo RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
      <LocationType>School</LocationType>
      <SiteCategory>Other</SiteCategory>
      <Name>Lincoln High School</Name>
      <LocalId>2343LHS</LocalId>
    </ChargedLocationInfo>
    
    
    

`
var test_CollectionAcquittal = `<CollectionAcquittal RefId="D3E34B35-9D75-101A-8C3D-00AA001A1653">
  <ReportingAuthorityList>
	<ReportingAuthority>
		<Name>St Joseph's Primary School, Leonville</Name>
		<System>Catholic</System>
		<AuthorityId>42732</AuthorityId>
	</ReportingAuthority>
	<ReportingAuthority>
		<Name>St Mary's Primary School, Leonville</Name>
		<System>Catholic</System>
		<AuthorityId>42733</AuthorityId>
	</ReportingAuthority>
  </ReportingAuthorityList>  
 <SubmittingAuthority>
		<Name>Paramatta Diocese</Name>
		<System>Catholic</System>
		<AuthorityId>423</AuthorityId>
  </SubmittingAuthority>
  <SubmittedBy>
    <Name>Fred Nurk</Name>
	<Role>General Manager</Role>
	<Organisation>Catholic Schools NSW</Organisation>
	<Signature>https://www.example.com/nurk.jpg</Signature>
	<SignatureImageType>01</SignatureImageType>
	<Date>2022-03-01</Date>
  </SubmittedBy>
    <AuditedBy>
    <Name>Jane Doe</Name>
	<Role>School Authority Auditor</Role>
	<Organisation>Doe &amp; Associates Accountants</Organisation>
	<Signature>https://www.example.com/doe.jpg</Signature>
	<SignatureImageType>01</SignatureImageType>
	<Date>2022-03-03</Date>
  </AuditedBy>
  <AuditorASICNumber>J83948349</AuditorASICNumber>
  <Recipient>NSW Education Department</Recipient>
  <Collection>STATS</Collection>
  <CollectionYear>2022</CollectionYear>
  <RoundCode>1</RoundCode>
  <Declaration>
  &lt;h1&gt;FINANCIAL ACCOUNTABILITY CERTIFICATION – 2022&lt;/h1&gt;
  
  &lt;p&gt;In receiving the NSW SRS Entitlement in respect of the NSW Government in 2022, on behalf of St Joseph's Primary School, Leonville and St Mary's Primary School, Leonville,
  Catholic Schools NSW agrees to spend, or commit to spend, these funds on any of the identified Educational Purposes listed in the letter.&lt;/p&gt;
  
  &lt;ol&gt;&lt;li&gt;I certify that the school/school authority does not and will not operate for profit within the meaning of Section 83C of the &lt;i&gt;Education Act 1990&lt;/i&gt;.&lt;/li&gt;
  &lt;li&gt;I certify that the school/school authority agrees to abide by any applicable guidelines or directions that may be issued from time to time pursuant to the 
  &lt;i&gt;Education Act 1990&lt;/i&gt; and &lt;i&gt;Education Regulation 2017&lt;/i&gt;.&lt;/li&gt;&lt;/ol&gt;
  </Declaration>
  <AuditorStatement>
  &lt;h1&gt;AUDITOR'S CERTIFICATION OF EXPENDITURE OF 2022 FUNDS AND 2022 ENROLMENTS&lt;/h1&gt;
  
  &lt;p&gt;Note:  Auditor must be registered with Australian Securities and Investments Commission (ASIC)&lt;/p&gt;

  &lt;ol&gt;&lt;li&gt;I certify to the best of my knowledge and have reasonable assurance that the NSW SRS Entitlement provided by the NSW Government in 2022
  was used in accordance with the Educational Purposes.&lt;/li&gt;
  &lt;li&gt;I certify to the best of my knowledge and have reasonable assurance that the enrolment figures provided under the 2022 National Schools Statistics Collection 
  are true and correct at the date of this certification.&lt;/li&gt;&lt;/ol&gt;
  </AuditorStatement>

</CollectionAcquittal>
`
var test_CollectionDeclaration = `<CollectionDeclaration RefId="D3E34B35-9D75-101A-8C3D-00AA001A1663">
  <ReportingAuthorityList>
	<ReportingAuthority>
		<Name>St Joseph's Primary School, Leonville</Name>
		<System>Catholic</System>
		<AuthorityId>42732</AuthorityId>
	</ReportingAuthority>
	<ReportingAuthority>
		<Name>St Mary's Primary School, Leonville</Name>
		<System>Catholic</System>
		<AuthorityId>42733</AuthorityId>
	</ReportingAuthority>
  </ReportingAuthorityList>  
 <SubmittingAuthority>
		<Name>Paramatta Diocese</Name>
		<System>Catholic</System>
		<AuthorityId>423</AuthorityId>
  </SubmittingAuthority>
  <SubmittedBy>
    <Name>Fred Nurk</Name>
	<Role>General Manager</Role>
	<Organisation>Catholic Schools NSW</Organisation>
	<Signature>https://www.example.com/nurk.jpg</Signature>
	<SignatureImageType>01</SignatureImageType>
	<Date>2022-03-01</Date>
  </SubmittedBy>
  <Recipient>NSW Education Department</Recipient>
  <Collection>STATS</Collection>
  <CollectionYear>2022</CollectionYear>
  <RoundCode>1</RoundCode>
  <Declaration>
  &lt;h1&gt;FINANCIAL ACCOUNTABILITY CERTIFICATION – 2022&lt;/h1&gt;
  
  &lt;p&gt;In receiving the NSW SRS Entitlement in respect of the NSW Government in 2022, on behalf of St Joseph's Primary School, Leonville and St Mary's Primary School, Leonville,
  Catholic Schools NSW agrees to spend, or commit to spend, these funds on any of the identified Educational Purposes listed in the letter.&lt;/p&gt;
  
   &lt;ol&gt;&lt;li&gt;I certify that the school/school authority does not and will not operate for profit within the meaning of Section 83C of the &lt;i&gt;Education Act 1990&lt;/i&gt;.&lt;/li&gt;
   &lt;li&gt;I certify that the school/school authority agrees to abide by any applicable guidelines or directions that may be issued from time to time pursuant to the 
   &lt;i&gt;Education Act 1990&lt;/i&gt; and &lt;i&gt;Education Regulation 2017&lt;/i&gt;.&lt;/li&gt;&lt;/ol&gt;
  </Declaration>

</CollectionDeclaration>
`
var test_CollectionRound = `<CollectionRound RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
  <AGCollection>STATS</AGCollection>
  <CollectionYear>2018</CollectionYear>
  <AGRoundList>
	<AGRound>
		<RoundCode>STATS1</RoundCode>
		<RoundName>Student Attendance Semester 1</RoundName>
		<StartDate>2019-06-28</StartDate>
		<DueDate>2019-08-02</DueDate>
		<EndDate>2019-10-01</EndDate>
	</AGRound>
	<AGRound>
		<RoundCode>STATS2 </RoundCode>
		<RoundName>Student Attendance Term 3</RoundName>
		<StartDate>2019-09-20</StartDate>
		<DueDate>2019-10-18</DueDate>
		<EndDate>2019-12-31</EndDate>
	</AGRound>
  </AGRoundList>  
</CollectionRound>
`
var test_CollectionStatus = `<CollectionStatus RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
  <ReportingAuthority>Ballarat Diocese</ReportingAuthority>
  <ReportingAuthoritySystem>Vic Catholic</ReportingAuthoritySystem>
  <ReportingAuthorityCommonwealthId>012345</ReportingAuthorityCommonwealthId>
  <SubmittedBy>XXXXXXXX</SubmittedBy>
  <SubmissionTimestamp>2018-05-30T09:00:00</SubmissionTimestamp>
  <AGCollection>FQ</AGCollection>
  <CollectionYear>2018</CollectionYear>
  <RoundCode>FQ1</RoundCode>
  <AGReportingObjectResponseList>		
	<AGReportingObjectResponse>
	  <SubmittedRefId>D3E34B35-9D75-101A-8C3D-00AA001A1652</SubmittedRefId>  
	  <SIFRefId>73048175-5738-4386-A260-40CE6230FC1B</SIFRefId>   
	  <HTTPStatusCode>202</HTTPStatusCode> 
      <CommonwealthId>012345</CommonwealthId>
      <EntityName>XXX Secondary College</EntityName>
	  <AGSubmissionStatusCode>Declared</AGSubmissionStatusCode>  
    </AGReportingObjectResponse>	
	<AGReportingObjectResponse>
	  <SubmittedRefId>E9C5106F-7E87-4773-ACCD-390F8E8B9314</SubmittedRefId>  
	  <HTTPStatusCode>400</HTTPStatusCode>
	  <ErrorText>Item components do not add up. Malformed characters.</ErrorText>
      <CommonwealthId>12387</CommonwealthId>
      <EntityName>ABCXYZ Secondary College</EntityName>
	  <AGSubmissionStatusCode>In Progress</AGSubmissionStatusCode>
      <AGRuleList>
	    <AGRule>
		  <AGRuleCode>WR-076</AGRuleCode>
			<AGRuleComment>Cannot be ignored because Components do not add up to Total - Fix</AGRuleComment>
			<AGRuleResponse></AGRuleResponse>  
			<AGRuleStatus>Fail</AGRuleStatus>
		</AGRule>
	  </AGRuleList>	
    </AGReportingObjectResponse>
	<AGReportingObjectResponse>
	  <SubmittedRefId>CFB3EB6B-F873-4108-BAA7-ED11711A9E86</SubmittedRefId>
	  <SIFRefId>4C194A64-362F-4163-BAE9-1A7852913E3C</SIFRefId>
	  <HTTPStatusCode>201</HTTPStatusCode>
      <CommonwealthId>12388</CommonwealthId>
      <EntityName>Ballarat Diocese</EntityName>
	  <AGSubmissionStatusCode>In Review</AGSubmissionStatusCode>
    </AGReportingObjectResponse>
  </AGReportingObjectResponseList>
</CollectionStatus>
`
var test_Debtor = `    <Debtor RefId="B5739375-800A-C4CC-6385-0BB2754114AA">
      <BilledEntity SIF_RefObject="StudentContactPersonal">8B231144-301C-4D3B-95D4-8BB74C866AE1</BilledEntity>
      <AddressList>
        <Address Type="0123" Role="2382">
          <Street>
            <Line1>23 E. 43rd St.</Line1>
          </Street>
          <City>Chicago</City>
          <StateProvince>IL</StateProvince>
          <Country>8104</Country>
          <PostalCode>60611</PostalCode>
        </Address>
      </AddressList>
      <BillingName>Flintstone Family Trust</BillingName>
      <BillingNote>Do not invoice under parent's real name.</BillingNote>
      <Discount>10.0</Discount>
	  <BSB>12121212</BSB>
	  <AccountNumber>34343434</AccountNumber>
	  <AccountName>Flintstone Inc.</AccountName>
    </Debtor>
    
`
var test_FinancialAccount = `    <FinancialAccount RefId="EEC8FC12-8D2C-4EE3-94A8-6C5395024EDE">
      <LocalId>89001</LocalId>
      <AccountNumber>9990001</AccountNumber>
      <Name>Purchased Foods</Name>
      <Description>Purchased Foods</Description>
      <ClassType>Asset</ClassType>
      <AccountCode>001-002-0034</AccountCode>
      <CreationDate>2003-01-01</CreationDate>
      <CreationTime>04:32:23-08:00</CreationTime> 
    </FinancialAccount>
`
var test_FinancialQuestionnaireCollection = `    <FinancialQuestionnaireCollection RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
		<FQYear>2018</FQYear> 
		<RoundCode>FQ1</RoundCode>
		<SoftwareVendorInfo>  
			<SoftwareProduct>MYOB</SoftwareProduct>
			<SoftwareVersion>11.0</SoftwareVersion>
		</SoftwareVendorInfo>
		<FQReportingList> 
			<FQReporting>
				<CommonwealthId>012345</CommonwealthId> 
				<EntityName>XXX Secondary College</EntityName>
				<EntityContact>
					<Name Type="LGL">
					  <FamilyName>Citizen</FamilyName> 
					  <GivenName>John</GivenName> 
					</Name>
					<PositionTitle>Auditor</PositionTitle> 
					<Email Type="01">jcitizen@school.vic.edu.au</Email> 
					<PhoneNumber Type="0096"><Number>03 1234-5678</Number></PhoneNumber> 
				</EntityContact>
				<FQContextualQuestionList>
					<FQContextualQuestion>
						<FQContext>GI.010</FQContext>
						<FQAnswer>Cash</FQAnswer>
					</FQContextualQuestion>
					<FQContextualQuestion>
						<FQContext>GI.015</FQContext>
						<FQAnswer>No</FQAnswer>
					</FQContextualQuestion>
				</FQContextualQuestionList>
				
				<FQItemList>
					<FQItem>
						<FQItemCode>RI.010</FQItemCode>
						<TuitionAmount>119827</TuitionAmount>
						<BoardingAmount>119827</BoardingAmount>
						<SystemAmount>12345</SystemAmount>
						<DioceseAmount>76543</DioceseAmount>
				    </FQItem>
					<FQItem>
						<FQItemCode>RI.020</FQItemCode>
						<TuitionAmount>119827</TuitionAmount>
						<BoardingAmount>119827</BoardingAmount>
						<SystemAmount>12345</SystemAmount>
						<DioceseAmount>76543</DioceseAmount>
				    </FQItem>
				    <FQItem>
						<FQItemCode>RI.030</FQItemCode>
						<TuitionAmount>119827</TuitionAmount>
						<BoardingAmount>119827</BoardingAmount>
						<SystemAmount>12345</SystemAmount>
						<DioceseAmount>76543</DioceseAmount>
					</FQItem>
				</FQItemList>
				<AGRuleList>
					<AGRule>
						<AGRuleCode>WR-076</AGRuleCode>
						<AGRuleComment>I ignored this rule because</AGRuleComment>
					</AGRule>
				</AGRuleList>
		  </FQReporting>
		  <FQReporting> 
				<CommonwealthId>12387</CommonwealthId>
				<EntityName>ABCXYZ Secondary College</EntityName>
				<EntityContact>
					<Name Type="LGL">
						 <FamilyName>Citizen</FamilyName>
						 <GivenName>John</GivenName>
					</Name>
					<PositionTitle>Auditor</PositionTitle>
					<Email Type="01">jcitizen@school.vic.edu.au</Email>
					<PhoneNumber Type="0096"><Number>03 1234-5678</Number></PhoneNumber>
				</EntityContact>
				<FQContextualQuestionList>
					<FQContextualQuestion>
						<FQContext>GI.010</FQContext>
						<FQAnswer>Cash</FQAnswer>
				    </FQContextualQuestion>
				    <FQContextualQuestion>
						<FQContext>GI.015</FQContext>
						<FQAnswer>No</FQAnswer>
					</FQContextualQuestion>
				</FQContextualQuestionList>
				<FQItemList>
					<FQItem>
						<FQItemCode>RI.010</FQItemCode>
						<TuitionAmount>119827</TuitionAmount>
						<BoardingAmount>119827</BoardingAmount>
						<SystemAmount>12345</SystemAmount>
						<DioceseAmount>76543</DioceseAmount>
				    </FQItem>
				    <FQItem>
						<FQItemCode>RI.020</FQItemCode>
						<TuitionAmount>119827</TuitionAmount>
						<BoardingAmount>119827</BoardingAmount>
						<SystemAmount>12345</SystemAmount>
						<DioceseAmount>76543</DioceseAmount>
				    </FQItem>
					<FQItem>
						<FQItemCode>RI.030</FQItemCode>
						<TuitionAmount>119827</TuitionAmount>
						<BoardingAmount>119827</BoardingAmount>
						<SystemAmount>12345</SystemAmount>
						<DioceseAmount>76543</DioceseAmount>
					</FQItem>
				</FQItemList>
				<AGRuleList>
				   <AGRule>
						<AGRuleCode>WR-076</AGRuleCode>
						<AGRuleComment>I ignored this rule because</AGRuleComment>
				   </AGRule>
				</AGRuleList>
			</FQReporting>
		    <FQReporting> 
				<CommonwealthId>12388</CommonwealthId>
				<EntityName>Ballarat Diocese</EntityName>
				<EntityContact>
					<Name Type="LGL">
					  <FamilyName>Citizen</FamilyName>
					  <GivenName>John</GivenName>
					</Name>
					<PositionTitle>Auditor</PositionTitle>
					<Email Type="01">jcitizeb@school.vic.edu.au</Email>
					<PhoneNumber Type="0096"><Number>03 1234-5678</Number></PhoneNumber>
				</EntityContact>
				<FQContextualQuestionList>
				  <FQContextualQuestion>
					<FQContext>GI.010</FQContext>
					<FQAnswer>Cash</FQAnswer>
				  </FQContextualQuestion>
				  <FQContextualQuestion>
						<FQContext>GI.015</FQContext>
						<FQAnswer>No</FQAnswer>
				  </FQContextualQuestion>
				</FQContextualQuestionList>
				<FQItemList>
				  <FQItem>
					<FQItemCode>RI.010</FQItemCode>
					<TuitionAmount>119827</TuitionAmount>
					<BoardingAmount>119827</BoardingAmount>
					<SystemAmount>12345</SystemAmount>
					<DioceseAmount>76543</DioceseAmount>
				  </FQItem>
				  <FQItem>
					<FQItemCode>RI.020</FQItemCode>
					<TuitionAmount>119827</TuitionAmount>
					<BoardingAmount>119827</BoardingAmount>
					<SystemAmount>12345</SystemAmount>
					<DioceseAmount>76543</DioceseAmount>
				  </FQItem>
				  <FQItem>
					<FQItemCode>RI.030</FQItemCode>
					<TuitionAmount>119827</TuitionAmount>
					<BoardingAmount>119827</BoardingAmount>
					<SystemAmount>12345</SystemAmount>
					<DioceseAmount>76543</DioceseAmount>
				  </FQItem>
				</FQItemList>
				<AGRuleList>
				  <AGRule>
					<AGRuleCode>WR-076</AGRuleCode>
					<AGRuleComment>I ignored this rule because</AGRuleComment>
				  </AGRule>
				</AGRuleList>
			  </FQReporting>
		  </FQReportingList>
		</FinancialQuestionnaireCollection>
`
var test_Invoice = `    <Invoice RefId="CA123458-47DE-A974-63FE-B238759FD123">
      <InvoicedEntity SIF_RefObject="Debtor">BA497254-965F-DA48-965A-BCE4589EA421</InvoicedEntity>
      <BillingDate>1999-10-12</BillingDate>
      <TransactionDescription>Activity Fees</TransactionDescription>
      <BilledAmount Type="Debit" Currency="AUD">50.00</BilledAmount> 
      <Ledger>Family</Ledger>
      <TaxRate>10.0</TaxRate>
      <TaxAmount Currency="AUD">5.00</TaxAmount>
      <CreatedBy>Peter Rabbit</CreatedBy>
      <ApprovedBy>Fred Flintstone</ApprovedBy>
      <ItemDetail>5 Widgets</ItemDetail>
      <FinancialAccountRefIdList>
        <FinancialAccountRefId>AE109F1A-C2DE-41E4-9DF5-C418F3DF18A8</FinancialAccountRefId>
      </FinancialAccountRefIdList>
      <AccountingPeriod>200004</AccountingPeriod>
    </Invoice>
`
var test_Journal = `    <Journal RefId="B5739375-800A-C4CC-6385-0BB2754114AA">
        <DebitFinancialAccountRefId>8B231144-301C-4D3B-95D4-8BB74C866AE1</DebitFinancialAccountRefId>
        <CreditFinancialAccountRefId>B7A34E56-1C97-345C-0A4E-11BB112B2753</CreditFinancialAccountRefId>
        <Amount Currency="AUD">7.00</Amount>
        <CreatedDate>2015-02-01</CreatedDate>
        <CreatedBy>Fred Flintstone</CreatedBy>
    </Journal>
    
`
var test_LibraryPatronStatus = `
    <LibraryPatronStatus RefId="D3E34B35-9D75-101A-8C3D-00AA001A1652">
	  <LibraryType>S</LibraryType>
	  <PatronRefId>5DE34B35-9D75-101A-8C3D-00AA001A1652</PatronRefId>
	  <PatronLocalId>MAR01</PatronLocalId>
	  <PatronRefObject>StudentPersonal</PatronRefObject>
      <ElectronicIdList>

        <ElectronicId Type="01">P 7676</ElectronicId>

        <ElectronicId Type="02">201906300007676JJI</ElectronicId>

        <ElectronicId Type="03">1234</ElectronicId>

        <ElectronicId Type="04">0EF3098055921FEDC008FEEB</ElectronicId>

      </ElectronicIdList>

      <TransactionList>

        <Transaction>

          <ItemInfo Type="LibraryMaterial">

            <Title>Redwall</Title>

            <Author>Brian Jacques</Author>

            <ElectronicId Type="01">T 311</ElectronicId>

            <CallNumber>YA JACQ</CallNumber>

            <Cost>6.99</Cost>
            
            <ReplacementCost>7.5</ReplacementCost>

          </ItemInfo>

          <CheckoutInfo>

			<CheckedOutOn>2018-12-04T23:59:59-05:00</CheckedOutOn> 	
            <ReturnBy>2019-01-04T23:59:59-05:00</ReturnBy>

          </CheckoutInfo>

          <FineInfoList>

            <FineInfo Type="Overdue">

              <Assessed>2019-01-06T13:19:00-05:00</Assessed>

              <Description>ESTIMATED FINE FOR OVERDUE ITEM</Description>

              <Amount>0.50</Amount>

            </FineInfo>

          </FineInfoList>

        </Transaction>

        <Transaction>

          <ItemInfo Type="LibraryMaterial">

            <Title>Mossflower</Title>

            <Author>Brian Jacques</Author>

            <ElectronicId Type="01">T 1211</ElectronicId>

            <CallNumber>YA JACQ</CallNumber>

            <Cost>7.49</Cost>

          </ItemInfo>

          <HoldInfoList>

            <HoldInfo Type="Ready">

              <DatePlaced>2018-12-18</DatePlaced>

              <DateNeeded>2019-01-16</DateNeeded>
              
              <ReservationExpiry>2019-01-31</ReservationExpiry>

              <MadeAvailable>2019-01-05</MadeAvailable>

              <Expires>2019-02-01</Expires>

            </HoldInfo>

          </HoldInfoList>

        </Transaction>

        <Transaction>

          <ItemInfo Type="LibraryMaterial">

            <Title>The Bellmaker</Title>

            <Author>Brian Jacques</Author>

          </ItemInfo>

          <HoldInfoList>

            <HoldInfo Type="NotReady">

              <DatePlaced>2018-12-18</DatePlaced>

            </HoldInfo>

          </HoldInfoList>

        </Transaction>

        <Transaction>

          <ItemInfo Type="Textbook">

            <Title>Passport to Algebra and Geometry</Title>

            <Author>Larson, Boswell, Kanold and Stiff</Author>

            <ElectronicId Type="01">98770000178215</ElectronicId>

            <Cost>57.18</Cost>

          </ItemInfo>

          <CheckoutInfo>
            <CheckedOutOn>2019-05-15T23:59:59-05:00</CheckedOutOn>

            <ReturnBy>2019-06-15T23:59:59-05:00</ReturnBy>

          </CheckoutInfo>

        </Transaction>

        <Transaction>

          <ItemInfo Type="LibraryMaterial">

            <Title>Harry Potter and the Goblet of Fire</Title>

            <Author>JK Rowlings</Author>

            <ElectronicId Type="01">T 99321</ElectronicId>

          </ItemInfo>

          <FineInfoList>

            <FineInfo Type="Refund">

              <Assessed>2019-01-02T09:29:00-05:00</Assessed>

              <Description>Lost item returned</Description>

              <Amount>28.90</Amount>

            </FineInfo>

          </FineInfoList>

        </Transaction>

        <Transaction>

          <FineInfoList>

            <FineInfo Type="Other">

              <Assessed>2019-01-04T10:14:00-05:00</Assessed>

              <Description>Running in library</Description>

              <Amount>1.00</Amount>

              <Reference>42</Reference>

            </FineInfo>

          </FineInfoList>

        </Transaction>

      </TransactionList>

      <MessageList>

        <Message Priority="Normal">

          <Text>This is a normal message.</Text>

        </Message>

        <Message Priority="Urgent">

          <Sent>2019-01-12T14:01:00-05:00</Sent>

          <Text>This is an URGENT message!</Text>

        </Message>

      </MessageList>

      <NumberOfCheckouts>1</NumberOfCheckouts>
      <NumberOfHoldItems>2</NumberOfHoldItems>

      <NumberOfOverdues>1</NumberOfOverdues>

      <NumberOfFines>1</NumberOfFines>

      <FineAmount>1.00</FineAmount>

      <NumberOfRefunds>1</NumberOfRefunds>

      <RefundAmount>28.90</RefundAmount>

    </LibraryPatronStatus>

`
var test_NAPCodeFrameexample1 = `<NAPCodeFrame RefId="09381381-641E-4034-A54B-9A84D7B54426">
  <NAPTestRefId>CE9FFC50-33B4-4F19-8DD5-62579B681A63</NAPTestRefId>
  <TestContent>
    <NAPTestLocalId>NAPLAN-2017-0001-Reading</NAPTestLocalId>
    <TestName>NAPLAN Reading Yr 3, 2017</TestName>
    <TestLevel><Code>3</Code></TestLevel>
    <TestType>Normal</TestType>
    <Domain>Reading</Domain>
    <TestYear>2017</TestYear>
    <StagesCount>3</StagesCount>
    <DomainBands>
      <Band1Lower>0</Band1Lower>
      <Band1Upper>258</Band1Upper>
      <Band2Lower>258</Band2Lower>
      <Band2Upper>318</Band2Upper>
      <Band3Lower>318</Band3Lower>
      <Band3Upper>368</Band3Upper>
      <Band4Lower>368</Band4Lower>
      <Band4Upper>417</Band4Upper>
      <Band5Lower>417</Band5Lower>
      <Band5Upper>466</Band5Upper>
      <Band6Lower>466</Band6Lower>
      <Band6Upper>526</Band6Upper>
      <Band7Lower>526</Band7Lower>
      <Band7Upper>574</Band7Upper>
      <Band8Lower>574</Band8Lower>
      <Band8Upper>633</Band8Upper>
      <Band9Lower>633</Band9Lower>
      <Band9Upper>681</Band9Upper>
      <Band10Lower>681</Band10Lower>
      <Band10Upper>999</Band10Upper>
    </DomainBands>
    <DomainProficiency>
       <Level1Lower>10</Level1Lower>
       <Level1Upper>20</Level1Upper>
       <Level2Lower>30</Level2Lower>
       <Level2Upper>40</Level2Upper>
       <Level3Lower>50</Level3Lower>
       <Level3Upper>60</Level3Upper>
       <Level4Lower>70</Level4Lower>
       <Level4Upper>80</Level4Upper>
    </DomainProficiency>
  </TestContent>
  <TestletList>  
    <Testlet>
      <NAPTestletRefId>74391492-B229-4B32-A9ED-E1B03C929604</NAPTestletRefId>
      <TestletContent>
        <NAPTestletLocalId>NAPLAN-2017-0001-Reading-A-00</NAPTestletLocalId>
        <TestletName>Testlet A-1 for Reading Yr 3</TestletName>
        <Node>A</Node>
        <LocationInStage>1</LocationInStage>
        <TestletMaximumScore>12</TestletMaximumScore>
      </TestletContent>
      <TestItemList>      
        <TestItem>
          <TestItemRefId>87734F2A-673B-4F2D-A805-532869CE3F57</TestItemRefId>
          <SequenceNumber>0</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-00</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>MC</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <MultipleChoiceOptionCount>4</MultipleChoiceOptionCount>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>kzgms666</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>70032978-57F3-4941-ADE7-BDC79ED34B5B</TestItemRefId>
          <SequenceNumber>1</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-01</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>MC</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <MultipleChoiceOptionCount>4</MultipleChoiceOptionCount>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>vzahj691</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>4CC4CB22-105E-485F-926C-BB1C8F0F4079</TestItemRefId>
          <SequenceNumber>2</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-02</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>SP</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>gmynn860</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>DD84E41B-59A9-41FF-B097-5D420DB661F1</TestItemRefId>
          <SequenceNumber>3</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-03</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>MC</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <MultipleChoiceOptionCount>4</MultipleChoiceOptionCount>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>zapvk881</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>E5D5FFFF-6168-4FC9-BC8E-A35E49799F19</TestItemRefId>
          <SequenceNumber>4</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-04</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>IO</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>kbrgs236</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>29B1B7C8-6D9D-4D84-BC2D-C88B6ED172DD</TestItemRefId>
          <SequenceNumber>5</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-05</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>MC</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <MultipleChoiceOptionCount>4</MultipleChoiceOptionCount>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>rokoh652</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>D9EDC321-D0F0-490A-8BE6-55A3C6548B60</TestItemRefId>
          <SequenceNumber>6</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-06</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>MC</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <MultipleChoiceOptionCount>4</MultipleChoiceOptionCount>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>qrlet247</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>F962D535-13F3-4837-9BFD-DCED4AEBEE19</TestItemRefId>
          <SequenceNumber>7</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-07</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>IGM</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>zhplu511</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>22B8B073-F6C6-40F3-8BF6-108FFA841E56</TestItemRefId>
          <SequenceNumber>8</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-08</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>MC</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <MultipleChoiceOptionCount>4</MultipleChoiceOptionCount>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>biyci381</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>AA59E605-D6C1-4332-8F47-60C40BC3BAC5</TestItemRefId>
          <SequenceNumber>9</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-09</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>TE</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AES</MarkingType>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>cbptn921</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>952B4AFC-36FE-4309-83C8-8EE398BA7F97</TestItemRefId>
          <SequenceNumber>10</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-10</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>PO</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>ioysv413</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
        <TestItem>
          <TestItemRefId>E138C292-FB89-4736-946D-16C78C8B3893</TestItemRefId>
          <SequenceNumber>11</SequenceNumber>
	  <TestItemContent>
            <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-11</NAPTestItemLocalId>
            <ItemName>Reading Unit A</ItemName>
            <ItemType>MC</ItemType>
    <Subdomain>Letters</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>true</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <MultipleChoiceOptionCount>4</MultipleChoiceOptionCount>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
    <StimulusList>    
        <Stimulus>
            <StimulusLocalId>odzlq067</StimulusLocalId>
            <TextGenre>Narrative</TextGenre>
            <TextType>Simple</TextType>
            <WordCount>300</WordCount>
            <TextDescriptor>A nice rollicking anecdote</TextDescriptor>
            <Content>http://example.com/Reading.xml</Content>
        </Stimulus>
    </StimulusList>
	  </TestItemContent>
        </TestItem>
      </TestItemList>
    </Testlet>
  </TestletList>
</NAPCodeFrame>
`
var test_NAPEventStudentLink = `    <NAPEventStudentLink RefId="2014F046-0C95-4613-8ACD-8A6CA8227C7A">
		  <StudentPersonalRefId>2F71D7C1-29B2-4FAD-9DB7-9FD80DC8804A</StudentPersonalRefId>
		  <PlatformStudentIdentifier>R100000004R</PlatformStudentIdentifier>
		  <SchoolInfoRefId>19F9C467-D299-4A35-9825-2148BF39DB64</SchoolInfoRefId>
		  <SchoolACARAId>21212</SchoolACARAId>
		  <NAPTestRefId>5195A2E9-37CE-4F51-AEF2-9095D9CF97FA</NAPTestRefId>
		  <NAPTestLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation</NAPTestLocalId>
		  <SchoolSector>NG</SchoolSector>
		  <System>0001</System>
		  <SchoolGeolocation>11</SchoolGeolocation>
		  <ReportingSchoolName>Podunk High School</ReportingSchoolName>
		  <NAPJurisdiction>1</NAPJurisdiction>
		  <ParticipationCode>P</ParticipationCode>
		  <ParticipationText>Present</ParticipationText>
		  <Device>iPad</Device>
		  <Date>2017-06-01</Date>
		  <StartTime>09:00:00</StartTime>
		  <LapsedTimeTest>PT2H57M32S</LapsedTimeTest>
		  <PersonalDetailsChanged>false</PersonalDetailsChanged>
		  <PSIOtherIdMatch>false</PSIOtherIdMatch>
		  <PossibleDuplicate>false</PossibleDuplicate>
		  <DOBRange>false</DOBRange>
   </NAPEventStudentLink>
`
var test_NAPStudentResponseSet = `  <NAPStudentResponseSet RefId="D90E7E40-33B2-44FA-9BA0-FAACA9D95487">
    <ReportExclusionFlag>false</ReportExclusionFlag>
    <CalibrationSampleFlag>false</CalibrationSampleFlag>
    <EquatingSampleFlag>false</EquatingSampleFlag>
    <PathTakenForDomain>A-D-F</PathTakenForDomain>
    <ParallelTest>A0-D1-F0</ParallelTest>
    <StudentPersonalRefId>2F71D7C1-29B2-4FAD-9DB7-9FD80DC8804A</StudentPersonalRefId>
    <PlatformStudentIdentifier>R100000004R</PlatformStudentIdentifier>
    <NAPTestRefId>615DD32A-A5A2-40A2-A417-187AE8F932EE</NAPTestRefId>
    <NAPTestLocalId>NAPLAN-2017-0001-Reading</NAPTestLocalId>
    <DomainScore>
        <RawScore>30.00</RawScore>
        <ScaledScoreValue>440.00</ScaledScoreValue>
        <ScaledScoreLogitValue>20.94</ScaledScoreLogitValue>
        <ScaledScoreStandardError>22.17</ScaledScoreStandardError>
        <ScaledScoreLogitStandardError>20.60</ScaledScoreLogitStandardError>
        <StudentDomainBand>6</StudentDomainBand>
        <StudentProficiency>Proficient</StudentProficiency>
        <PlausibleScaledValueList>
          <PlausibleScaledValue>14</PlausibleScaledValue>
          <PlausibleScaledValue>15</PlausibleScaledValue>
          <PlausibleScaledValue>16</PlausibleScaledValue>
          <PlausibleScaledValue>17</PlausibleScaledValue>
          <PlausibleScaledValue>18</PlausibleScaledValue>
        </PlausibleScaledValueList>
    </DomainScore>
    <TestletList>

        <Testlet>
            <NAPTestletRefId>1FD3A68E-530D-491F-920D-77F77074D159</NAPTestletRefId>
            <NAPTestletLocalId>NAPLAN-2017-0001-Reading-A-00</NAPTestletLocalId>
            <TestletSubScore>11</TestletSubScore>
            <ItemResponseList>
                <ItemResponse>
                    <NAPTestItemRefId>D9660081-EC37-4E20-B8FD-0F20896BA338</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-00</NAPTestItemLocalId>
                    <Response>A</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>1</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>04F23B44-17E0-4D34-B321-C2ED8C477D43</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-01</NAPTestItemLocalId>
                    <Response>C</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>2</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>F51E2743-62B0-491A-BE34-B3DC722E4A1F</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-02</NAPTestItemLocalId>
                    <Response>gKqvdfwIYE</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>3</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>859CDE35-4236-47C2-B593-813084D3E67B</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-03</NAPTestItemLocalId>
                    <Response>A</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>4</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>5A45B1B5-05B5-429A-9DA2-9757CCD18B99</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-04</NAPTestItemLocalId>
                    <Response>isRnrsoOzv</Response>
                    <ResponseCorrectness>Incorrect</ResponseCorrectness>
                    <Score>0</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>5</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>FC1281E6-0CF0-4E01-B160-C9D83EDB1F97</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-05</NAPTestItemLocalId>
                    <Response>A</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>6</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>A7F24BCE-4663-40CB-8587-24135B90345A</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-06</NAPTestItemLocalId>
                    <Response>C</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>7</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>7ABC7320-56C2-401E-B287-EFD7D1E90E12</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-07</NAPTestItemLocalId>
                    <Response>D</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>8</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>B2BC0B19-E52A-46C9-BD40-CD4576756746</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-08</NAPTestItemLocalId>
                    <Response>B</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>9</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>ECF360FB-F7D6-4DB7-B76B-AA598E4C4981</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-09</NAPTestItemLocalId>
                    <Response>B</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>10</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>EDCC84CA-78D7-41CC-9C82-7E5DEA8F35AA</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-10</NAPTestItemLocalId>
                    <Response>aFdOmmNEcv</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>11</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>7786117B-2C8E-40D4-A38E-EE3858FCE1B7</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-A-00-11</NAPTestItemLocalId>
                    <Response>iIauaAgmYC</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>12</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
            </ItemResponseList>
        </Testlet>
        <Testlet>
            <NAPTestletRefId>953CDBD8-B5F2-4703-A4AB-3B394A33C228</NAPTestletRefId>
            <NAPTestletLocalId>NAPLAN-2017-0001-Reading-D-01</NAPTestletLocalId>
            <TestletSubScore>11</TestletSubScore>
            <ItemResponseList>
                <ItemResponse>
                    <NAPTestItemRefId>8B89C8B9-6C05-42BA-A9D1-321B79A34241</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-00</NAPTestItemLocalId>
                    <Response>ZENYKCkmWr</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>1</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>854F56CE-C30F-49D3-AD15-9774D5113519</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-01</NAPTestItemLocalId>
                    <Response>C</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>2</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>7258B5BD-38E8-41A2-B27E-A89624DBCEC1</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-02</NAPTestItemLocalId>
                    <Response>A</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>3</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>BB94CCDD-21BD-43B6-B252-B5F628C32F80</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-03</NAPTestItemLocalId>
                    <Response>A</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>4</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>BD1F3E74-9362-4094-8902-FD865F76CB70</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-04</NAPTestItemLocalId>
                    <Response>D</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>5</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>D3D22E43-C431-48F5-BAD6-723C923C864C</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-05</NAPTestItemLocalId>
                    <Response>A</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>6</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>A5663CC8-27E3-4830-926C-3349465DAEE3</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-06</NAPTestItemLocalId>
                    <Response>C</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>7</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>DCA323A6-3B29-4942-A979-C594E363277D</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-07</NAPTestItemLocalId>
                    <Response>NtTIYMmIMC</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>8</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>35B98745-B527-4662-8891-CA12BED0E480</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-08</NAPTestItemLocalId>
                    <Response>fFnySwRyen</Response>
                    <ResponseCorrectness>Incorrect</ResponseCorrectness>
                    <Score>0</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>9</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>34A77184-276A-4C8B-BDA7-8E9D5446C74A</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-09</NAPTestItemLocalId>
                    <Response>D</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>10</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>B1E5EE81-6E58-4F71-9583-F1DDDE235892</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-10</NAPTestItemLocalId>
                    <Response>OlHtaQNHoM</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>11</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>ED1E5B13-782C-4BF7-8E83-3F0893C2F856</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-D-01-11</NAPTestItemLocalId>
                    <Response>C</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>12</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
            </ItemResponseList>
        </Testlet>
        <Testlet>
            <NAPTestletRefId>CC94B00C-5AAE-4B24-A58E-A4C4B04D0FD4</NAPTestletRefId>
            <NAPTestletLocalId>NAPLAN-2017-0001-Reading-F-00</NAPTestletLocalId>
            <TestletSubScore>8</TestletSubScore>
            <ItemResponseList>
                <ItemResponse>
                    <NAPTestItemRefId>D9AC94C9-461B-4E6D-9005-FCE134C5BF3E</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-00</NAPTestItemLocalId>
                    <Response>B</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>1</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>B4603868-3026-4043-A1E9-83C1D663392D</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-01</NAPTestItemLocalId>
                    <Response>cFGrURFLQJ</Response>
                    <ResponseCorrectness>Incorrect</ResponseCorrectness>
                    <Score>0</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>2</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>FB38C594-F4E4-4469-BBE5-4AF302171075</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-02</NAPTestItemLocalId>
                    <Response>wXMrYIUEef</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>3</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>35732438-2182-45B1-9059-04FCB3567637</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-03</NAPTestItemLocalId>
                    <Response>hZtAellWwR</Response>
                    <ResponseCorrectness>Incorrect</ResponseCorrectness>
                    <Score>0</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>4</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>602ECE01-57B0-48ED-90A5-070DAC7ED6CC</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-04</NAPTestItemLocalId>
                    <Response>ufSagjiuzw</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>5</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>8B7CCBE1-E9CB-4AAC-97F1-514DCEF41CDA</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-05</NAPTestItemLocalId>
                    <Response>eAnZrOBhEV</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>6</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>FD382A1F-D98D-4202-8DEE-99D11B230E78</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-06</NAPTestItemLocalId>
                    <Response>GlMlPTIYIO</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>7</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>316EDF8E-9E96-4979-8B3A-EE3B3B0A5258</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-07</NAPTestItemLocalId>
                    <Response>WHwIVewqtV</Response>
                    <ResponseCorrectness>Incorrect</ResponseCorrectness>
                    <Score>0</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>8</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>354069AF-929D-431B-A297-D3B78E356476</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-08</NAPTestItemLocalId>
                    <Response>cAIJHatCZI</Response>
                    <ResponseCorrectness>NotInPath</ResponseCorrectness>
                    <Score>0</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>9</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>31CF0277-5AED-4D0D-BFC5-2235F7B4F674</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-09</NAPTestItemLocalId>
                    <Response>ZcCCskdAXf</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>10</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>D91F3243-134C-4BA4-9F4B-E02D65671401</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-10</NAPTestItemLocalId>
                    <Response>nKImzEMCpy</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>11</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
                <ItemResponse>
                    <NAPTestItemRefId>EB7B780E-2C64-430E-9F44-B52DCD273E04</NAPTestItemRefId>
                    <NAPTestItemLocalId>NAPLAN-2017-0001-Reading-F-00-11</NAPTestItemLocalId>
                    <Response>B</Response>
                    <ResponseCorrectness>Correct</ResponseCorrectness>
                    <Score>1</Score>
                    <LapsedTimeItem>PT50S</LapsedTimeItem>
                    <SequenceNumber>12</SequenceNumber>
                    <ItemWeight>1</ItemWeight>
                </ItemResponse>        
            </ItemResponseList>
        </Testlet>
    </TestletList>
</NAPStudentResponseSet>    
`
var test_NAPTest = ` <NAPTest RefId="27BD7B5B-19F2-48A9-B645-FD6569E7654E">
  <TestContent>
    <NAPTestLocalId>NAPLAN-2017-0001-Reading</NAPTestLocalId>
    <TestName>NAPLAN Reading Yr 3, 2017</TestName>
    <TestLevel><Code>3</Code></TestLevel>
    <TestType>Normal</TestType>
    <Domain>Reading</Domain>
    <TestYear>2017</TestYear>
    <StagesCount>3</StagesCount>
    <DomainBands>
      <Band1Lower>0</Band1Lower>
      <Band1Upper>258</Band1Upper>
      <Band2Lower>258</Band2Lower>
      <Band2Upper>318</Band2Upper>
      <Band3Lower>318</Band3Lower>
      <Band3Upper>368</Band3Upper>
      <Band4Lower>368</Band4Lower>
      <Band4Upper>417</Band4Upper>
      <Band5Lower>417</Band5Lower>
      <Band5Upper>466</Band5Upper>
      <Band6Lower>466</Band6Lower>
      <Band6Upper>526</Band6Upper>
      <Band7Lower>526</Band7Lower>
      <Band7Upper>574</Band7Upper>
      <Band8Lower>574</Band8Lower>
      <Band8Upper>633</Band8Upper>
      <Band9Lower>633</Band9Lower>
      <Band9Upper>681</Band9Upper>
      <Band10Lower>681</Band10Lower>
      <Band10Upper>999</Band10Upper>
    </DomainBands>
    <DomainProficiency>
       <Level1Lower>10</Level1Lower>
       <Level1Upper>20</Level1Upper>
       <Level2Lower>30</Level2Lower>
       <Level2Upper>40</Level2Upper>
       <Level3Lower>50</Level3Lower>
       <Level3Upper>60</Level3Upper>
       <Level4Lower>70</Level4Lower>
       <Level4Upper>80</Level4Upper>
    </DomainProficiency>
  </TestContent>
</NAPTest>
`
var test_NAPTestItem = `<NAPTestItem RefId="31CA1A02-4E53-4A29-AE1E-807489C30827">
  <TestItemContent>
    <NAPTestItemLocalId>NAPLAN-2017-0005-Language Conventions: Spelling-S2-01-01</NAPTestItemLocalId>
    <ItemName>Spelling Unit S2</ItemName>
    <ItemType>MC</ItemType>
    <Subdomain>Spelling</Subdomain>
    <ItemDescriptor>Descriptor #1</ItemDescriptor>
    <ReleasedStatus>false</ReleasedStatus>
    <MarkingType>AS</MarkingType>
    <MultipleChoiceOptionCount>4</MultipleChoiceOptionCount>
    <CorrectAnswer>7</CorrectAnswer>
    <MaximumScore>1</MaximumScore>
    <ItemDifficulty>3</ItemDifficulty>
    <ItemDifficultyLogit5>.8</ItemDifficultyLogit5>
    <ItemDifficultyLogit62>.9</ItemDifficultyLogit62>
    <ItemDifficultyLogit5SE>.8</ItemDifficultyLogit5SE>
    <ItemDifficultyLogit62SE>.9</ItemDifficultyLogit62SE>
    <ItemProficiencyBand>3</ItemProficiencyBand>
    <ItemProficiencyLevel>C</ItemProficiencyLevel>
    <ExemplarURL>http://example.com/n3.xml</ExemplarURL>
    <ContentDescriptionList>
        <ContentDescription>MNA32</ContentDescription>
        <ContentDescription>MNA37</ContentDescription>
    </ContentDescriptionList>
  </TestItemContent>
</NAPTestItem>

`
var test_NAPTestScoreSummary = `<NAPTestScoreSummary RefId="E2721E52-7661-4A51-BF16-B4CAF12DC2F8">
  <SchoolInfoRefId>A35CE161-BFDB-4460-952B-7F87789A739C</SchoolInfoRefId>
  <SchoolACARAId>21212</SchoolACARAId>
  <NAPTestRefId>FBB8A444-1891-4701-B306-F3CED219369C</NAPTestRefId>
  <NAPTestLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation</NAPTestLocalId>
  <DomainNationalAverage>399</DomainNationalAverage>
  <DomainSchoolAverage>404</DomainSchoolAverage>
  <DomainJurisdictionAverage>402</DomainJurisdictionAverage>
  <DomainTopNational60Percent>416</DomainTopNational60Percent>
  <DomainBottomNational60Percent>376</DomainBottomNational60Percent>
</NAPTestScoreSummary>
`
var test_NAPTestlet = ` <NAPTestlet RefId="42DE5F3F-DA32-4DA6-8966-4C8FA176E4B4">
  <NAPTestRefId>FBB8A444-1891-4701-B306-F3CED219369C</NAPTestRefId>
  <NAPTestLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation</NAPTestLocalId>
  <TestletContent>
    <NAPTestletLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00</NAPTestletLocalId>
    <TestletName>Testlet C-1 for Language Conventions: Grammar and Punctuation Yr 3</TestletName>
    <Node>C</Node>
    <LocationInStage>1</LocationInStage>
    <TestletMaximumScore>12</TestletMaximumScore>
  </TestletContent>
  <TestItemList>
      <TestItem>
        <TestItemRefId>C77240DF-6E21-4774-AD1D-9454D5F60758</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-00</TestItemLocalId>
        <SequenceNumber>1</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>1C4016D1-AAAA-4034-BE8A-0CDE847B2EC7</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-01</TestItemLocalId>
        <SequenceNumber>2</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>3CCD15F4-BAF5-4192-A361-3C4184AB4FC8</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-02</TestItemLocalId>
        <SequenceNumber>3</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>66602398-5EC9-46E4-BCE6-E8917F49EE56</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-03</TestItemLocalId>
        <SequenceNumber>4</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>D1A5A9F2-43FB-40BB-A88F-71A879C5450D</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-04</TestItemLocalId>
        <SequenceNumber>5</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>3D253985-4021-418B-BD20-992EEA886535</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-05</TestItemLocalId>
        <SequenceNumber>6</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>D5B39AE5-B134-4909-A581-F69205F2682A</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-06</TestItemLocalId>
        <SequenceNumber>7</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>E5537981-2898-46DE-BC1C-3CEC4AFF9468</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-07</TestItemLocalId>
        <SequenceNumber>8</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>900CC548-23FE-49B5-ADCD-25854FBFA995</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-08</TestItemLocalId>
        <SequenceNumber>9</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>476FE04D-81CA-4C08-A555-60463F0DC918</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-09</TestItemLocalId>
        <SequenceNumber>10</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>217B9585-2D66-4270-ABD4-2AD57E7453B5</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-10</TestItemLocalId>
        <SequenceNumber>11</SequenceNumber>
      </TestItem>
      <TestItem>
        <TestItemRefId>982D24CB-A976-4AD1-B9D5-5CD4C09F8C00</TestItemRefId>
        <TestItemLocalId>NAPLAN-2017-0004-Language Conventions: Grammar and Punctuation-C-00-11</TestItemLocalId>
        <SequenceNumber>12</SequenceNumber>
      </TestItem>
  </TestItemList>
</NAPTestlet>

`
var test_PaymentReceipt = `    <PaymentReceipt RefId="EDF4985A-BC34-58FE-DA75-493AB3812345">
      <TransactionType>GL payment</TransactionType>
      <InvoiceRefId>CA123458-47DE-A974-63FE-B238759FD123</InvoiceRefId>
       <TransactionDate>1999-10-20</TransactionDate>
      <TransactionAmount Currency="AUD" Type="Debit">50.00</TransactionAmount>
      <ReceivedTransactionId>XYZZY</ReceivedTransactionId>
      <FinancialAccountRefIdList>
        <FinancialAccountRefId>ED12345F-DA84-9727-5BC2-8AA349DD3721</FinancialAccountRefId> 
      </FinancialAccountRefIdList>
    </PaymentReceipt>
    
`
var test_example0218 = `    <PurchaseOrder RefId="ED12345F-DA84-9727-5BC2-8AA349DD3721">
      <FormNumber>00342</FormNumber>
      <VendorInfoRefId>BD12345F-DA84-9727-5BC2-8AA349DD3723</VendorInfoRefId>
      <ChargedLocationInfoRefId>ED12345F-DA84-9727-5BC2-8AA349DD3722</ChargedLocationInfoRefId>
      <EmployeePersonalRefId>AD12345F-DA84-9727-5BC2-8AA349DD3721</EmployeePersonalRefId>
      <PurchasingItems>
        <PurchasingItem>
          <ItemNumber>154486</ItemNumber>
          <ItemDescription>Floor Wax</ItemDescription>
          <Quantity>10</Quantity>
          <UnitCost Currency="AUD">7.00</UnitCost>
          <ExpenseAccounts>
            <ExpenseAccount>
              <AccountCode>10-1100-610</AccountCode>
              <Amount Currency="AUD">70.00</Amount>
            </ExpenseAccount>
          </ExpenseAccounts>
        </PurchasingItem>
      </PurchasingItems> 
    </PurchaseOrder>
`
var test_StudentAttendanceCollectionExample2 = `<StudentAttendanceCollection RefId="D3E34B35-9D75-101A-8C3D-00AA001A1653">
  <StudentAttendanceCollectionYear>2020</StudentAttendanceCollectionYear>
  
  <RoundCode>Semester1</RoundCode>
  <SoftwareVendorInfo>
    
    <SoftwareProduct>Software</SoftwareProduct>
    <SoftwareVersion>Websys 1.0</SoftwareVersion>
  </SoftwareVendorInfo>
  <StudentAttendanceCollectionReportingList>
    <StudentAttendanceCollectionReporting>
      <CommonwealthId>24680</CommonwealthId>
      
      <EntityName>XXX Secondary College</EntityName>
      <EntityContact>
        <Name Type="LGL">
          <FamilyName>Citizen</FamilyName>
          
          <GivenName>John</GivenName>
          
        </Name>
        <PositionTitle>BusinessManager</PositionTitle>
        
       <Email Type="01">jcitizen@school.vic.edu.au</Email>
        
        <PhoneNumber Type="0096">
          <Number>03 1234-5678</Number>
        </PhoneNumber>
        
      </EntityContact>
      <StatsCohortYearLevelList>
        <StatsCohortYearLevel>
          <CohortYearLevel>
            <Code>1</Code>
          </CohortYearLevel>
          <StatsCohortList>
            <StatsCohort>
              <StatsCohortId>1000</StatsCohortId>
              <StatsIndigenousStudentType>T</StatsIndigenousStudentType>
              <CohortGender>M</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>3600</PossibleSchoolDays>
              <AttendanceDays>3400</AttendanceDays>
              <AttendanceLess90Percent>10</AttendanceLess90Percent>
              <AttendanceGTE90Percent>20</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>2400</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1001</StatsCohortId>
              <StatsIndigenousStudentType>T</StatsIndigenousStudentType>
              <CohortGender>F</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>3600</PossibleSchoolDays>
              <AttendanceDays>3300</AttendanceDays>
              <AttendanceLess90Percent>10</AttendanceLess90Percent>
              <AttendanceGTE90Percent>20</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>2400</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1002</StatsCohortId>
              <StatsIndigenousStudentType>T</StatsIndigenousStudentType>
              <CohortGender>X</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>120</PossibleSchoolDays>
              <AttendanceDays>118</AttendanceDays>
              <AttendanceLess90Percent>0</AttendanceLess90Percent>
              <AttendanceGTE90Percent>1</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>120</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1003</StatsCohortId>
              <StatsIndigenousStudentType>I</StatsIndigenousStudentType>
              <CohortGender>M</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>240</PossibleSchoolDays>
              <AttendanceDays>235</AttendanceDays>
              <AttendanceLess90Percent>0</AttendanceLess90Percent>
              <AttendanceGTE90Percent>2</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>240</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1004</StatsCohortId>
              <StatsIndigenousStudentType>I</StatsIndigenousStudentType>
              <CohortGender>F</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>120</PossibleSchoolDays>
              <AttendanceDays>120</AttendanceDays>
              <AttendanceLess90Percent>0</AttendanceLess90Percent>
              <AttendanceGTE90Percent>1</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>120</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1005</StatsCohortId>
              <StatsIndigenousStudentType>I</StatsIndigenousStudentType>
              <CohortGender>X</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>120</PossibleSchoolDays>
              <AttendanceDays>80</AttendanceDays>
              <AttendanceLess90Percent>1</AttendanceLess90Percent>
              <AttendanceGTE90Percent>0</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>0</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
          </StatsCohortList>
        </StatsCohortYearLevel>
        <StatsCohortYearLevel>
          <CohortYearLevel>
            <Code>2</Code>
          </CohortYearLevel>
          <StatsCohortList>
            <StatsCohort>
              <StatsCohortId>1006</StatsCohortId>
              <StatsIndigenousStudentType>T</StatsIndigenousStudentType>
              <CohortGender>M</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>1200</PossibleSchoolDays>
              <AttendanceDays>1200</AttendanceDays>
              <AttendanceLess90Percent>10</AttendanceLess90Percent>
              <AttendanceGTE90Percent>20</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>2400</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1007</StatsCohortId>
              <StatsIndigenousStudentType>T</StatsIndigenousStudentType>
              <CohortGender>F</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>3600</PossibleSchoolDays>
              <AttendanceDays>120</AttendanceDays>
              <AttendanceLess90Percent>10</AttendanceLess90Percent>
              <AttendanceGTE90Percent>20</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>2400</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1008</StatsCohortId>
              <StatsIndigenousStudentType>T</StatsIndigenousStudentType>
              <CohortGender>X</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>100</PossibleSchoolDays>
              <AttendanceDays>99</AttendanceDays>
              <AttendanceLess90Percent>0</AttendanceLess90Percent>
              <AttendanceGTE90Percent>1</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>100</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1009</StatsCohortId>
              <StatsIndigenousStudentType>I</StatsIndigenousStudentType>
              <CohortGender>M</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>600</PossibleSchoolDays>
              <AttendanceDays>580</AttendanceDays>
              <AttendanceLess90Percent>1</AttendanceLess90Percent>
              <AttendanceGTE90Percent>4</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>480</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1010</StatsCohortId>
              <StatsIndigenousStudentType>I</StatsIndigenousStudentType>
              <CohortGender>F</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>720</PossibleSchoolDays>
              <AttendanceDays>710</AttendanceDays>
              <AttendanceLess90Percent>0</AttendanceLess90Percent>
              <AttendanceGTE90Percent>6</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>720</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
            <StatsCohort>
              <StatsCohortId>1011</StatsCohortId>
              <StatsIndigenousStudentType>I</StatsIndigenousStudentType>
              <CohortGender>X</CohortGender>
              <DaysInReferencePeriod>120</DaysInReferencePeriod>
              <PossibleSchoolDays>120</PossibleSchoolDays>
              <AttendanceDays>120</AttendanceDays>
              <AttendanceLess90Percent>0</AttendanceLess90Percent>
              <AttendanceGTE90Percent>1</AttendanceGTE90Percent>
              <PossibleSchoolDaysGT90PercentAttendance>120</PossibleSchoolDaysGT90PercentAttendance>
            </StatsCohort>
          </StatsCohortList>
        </StatsCohortYearLevel>
      </StatsCohortYearLevelList>
    </StudentAttendanceCollectionReporting>
  </StudentAttendanceCollectionReportingList>
</StudentAttendanceCollection>
`
var test_StudentDataTransferNoteExample = `<StudentDataTransferNote RefId="D3E34B35-9D75-101A-8C3D-00AA001A1654">
  <Name Type="LGL">
    <FamilyName>Smith</FamilyName>
    <GivenName>Fred</GivenName>
    <MiddleName>Ebenezer</MiddleName>
	<PreferredFamilyName>Smith</PreferredFamilyName>
	<PreferredGivenName>Freddy</PreferredGivenName>
  </Name>
  <Gender>1</Gender>
  <BirthDate>2010-03-01</BirthDate>
  <PlaceOfBirth>Clayton</PlaceOfBirth>
  <StateOfBirth>VIC</StateOfBirth>
  <CountryOfBirth>1101</CountryOfBirth>
 <CountriesOfCitizenship>
    <CountryOfCitizenship>8104</CountryOfCitizenship>
    <CountryOfCitizenship>1101</CountryOfCitizenship>
  </CountriesOfCitizenship>
  <ArrivalSchool>
    <CommonwealthId>12345</CommonwealthId>
	<Name>Buchanan Secondary College</Name>
  </ArrivalSchool>
  <DepartureSchool>
    <ACARAId>30000</ACARAId>
	<Name>Lincoln Secondary College</Name>
	<City>Lincoln</City>
  <SchoolContactList>
    <SchoolContact>
      <PublishInDirectory>Y</PublishInDirectory>
      <ContactInfo>
        <Name Type="LGL">
          <Title>Mr</Title>
          <FamilyName>Miller</FamilyName>
          <GivenName>James</GivenName>
          <MiddleName>Mark</MiddleName>
          <Suffix>Jr.</Suffix>
          <FullName>Mr James Mark Miller Jr.</FullName>
        </Name>
        <PositionTitle>Business Manager</PositionTitle>
        <Role>School Information Contact Point</Role>
        <Address Type="0123" Role="012B">
          <Street>
            <Line1>23 Nicholson Street</Line1>
          </Street>
          <City>Carnegie</City>
          <StateProvince>VIC</StateProvince>
          <Country>1101</Country>
          <PostalCode>3004</PostalCode>
          <GridLocation>
            <Latitude>23.9876</Latitude>
            <Longitude>-98.8765</Longitude>
          </GridLocation>
        </Address>
        <EmailList>
          <Email Type="01">jmiller@lsc.vic.edu.au</Email>
          <Email Type="02">jmiller@yahoo.com.au</Email>
        </EmailList>
        <PhoneNumberList>
          <PhoneNumber Type="0096">
            <Number>03 9637-2000</Number>
            <Extension>72345</Extension>
            <ListedStatus>Y</ListedStatus>
          </PhoneNumber>
        </PhoneNumberList>
      </ContactInfo>
    </SchoolContact>
  </SchoolContactList> 
  </DepartureSchool>
  <PreviousSchoolList>
  <PreviousSchool>
    <ACARAId>30001</ACARAId>
	<Name>Lincoln Primary School</Name>
	<City>Lincolnville</City>
  </PreviousSchool> 
  <PreviousSchool>
    <ACARAId>30002</ACARAId>
	<Name>Buchanan Primary School</Name>
	<City>Lincolnville</City>
  </PreviousSchool> 
  </PreviousSchoolList>
  <NAPLANScoreList>
    <NAPLANScore>
	  <Domain>Grammar and Punctuation</Domain>
	  <ParticipationCode>P</ParticipationCode>
 <DomainScore>
    <ScaledScoreValue>440.00</ScaledScoreValue>
    <StudentDomainBand>6</StudentDomainBand>
    <StudentProficiency>Proficient</StudentProficiency>
  </DomainScore>
    <TestLevel><Code>3</Code></TestLevel>
	<TestYear>2017</TestYear>
	</NAPLANScore>
     <NAPLANScore>
	  <Domain>Numeracy</Domain>
	  <ParticipationCode>E</ParticipationCode>
	</NAPLANScore>
    <NAPLANScore>
	  <Domain>Reading</Domain>
	  <ParticipationCode>P</ParticipationCode>
<DomainScore>
    <ScaledScoreValue>440.00</ScaledScoreValue>
    <StudentDomainBand>6</StudentDomainBand>
    <StudentProficiency>Proficient</StudentProficiency>
  </DomainScore>
	</NAPLANScore>
    <NAPLANScore>
	  <Domain>Spelling</Domain>
	  <ParticipationCode>P</ParticipationCode>
<DomainScore>
    <ScaledScoreValue>440.00</ScaledScoreValue>
    <StudentDomainBand>6</StudentDomainBand>
    <StudentProficiency>Proficient</StudentProficiency>
  </DomainScore>
	</NAPLANScore>
    <NAPLANScore>
	  <Domain>Writing</Domain>
	  <ParticipationCode>X</ParticipationCode>
	</NAPLANScore>
 </NAPLANScoreList>
 <NCCDList>
   <NCCD>
     <LevelOfAdjustment>QDTP</LevelOfAdjustment>
	 <CategoryOfDisability>Physical</CategoryOfDisability>
	 <DisabilityCategoryConsideredList>
		<DisabilityCategoryConsidered>Physical</DisabilityCategoryConsidered>
		<DisabilityCategoryConsidered>Social-Emotional</DisabilityCategoryConsidered>
	 </DisabilityCategoryConsideredList>
	 <DateOfAssessment>2020-03-02</DateOfAssessment>
   </NCCD>
  </NCCDList>
 <FollowupRequest>true</FollowupRequest>
 <ChildSubjectToOrders>true</ChildSubjectToOrders>
 <Attendance>false</Attendance>
 <NationalUniqueStudentIdentifier>0001001</NationalUniqueStudentIdentifier>
 <YearLevel>
    <Code>10</Code>
  </YearLevel>
  <IndigenousStatus>1</IndigenousStatus>
  <LBOTE>Y</LBOTE>
<VisaStatus>
    <Code>101</Code>
	<VisaExpiryDate>2025-12-31</VisaExpiryDate>
 </VisaStatus>
   <OtherNames>
    <Name Type="AKA">
      <FamilyName>Anderson</FamilyName>
      <GivenName>Samuel</GivenName>
      <FullName>Samuel Anderson</FullName>
    </Name>
    <Name Type="PRF">
      <FamilyName>Rowinski</FamilyName>
      <GivenName>Sam</GivenName>
      <FullName>Sam Rowinski</FullName>
    </Name>
  </OtherNames>
  <EducationalAssessmentList>
    <EducationalAssessment>
	  <Name>PAT</Name>
	  <Content>PAT scores for student...</Content>
    </EducationalAssessment>
     <EducationalAssessment>
	  <Name>SAT</Name>
	  <Content>SAT scores for student...</Content>
    </EducationalAssessment>
  </EducationalAssessmentList>
  <StudentGradeList>
    <StudentGrade>
	  <Subject>Australian History: Gold and Iron</Subject>
	  <LearningArea>
	    <ACStrand>H</ACStrand>
	  <SubjectArea><Code>Australian History</Code></SubjectArea>
	  </LearningArea>
	  <Grade>
	    <Letter>B</Letter>
		<Narrative>Could do better</Narrative>
	  </Grade>
	</StudentGrade>
    <StudentGrade>
	  <Subject>World History: Revolutions</Subject>
	  <LearningArea>
	    <ACStrand>H</ACStrand>
	  <SubjectArea><Code>World History</Code></SubjectArea>
	  </LearningArea>
	  <Grade>
	    <Letter>A</Letter>
		<Narrative>Could not do better</Narrative>
	  </Grade>
	</StudentGrade>
  </StudentGradeList>
 </StudentDataTransferNote>
`
var test_VendorInfo = `    <VendorInfo RefId="AB3647C5-6865-4CF4-5678-DD34EF564E22">
      <Name>ABC School Supply</Name>
      <ContactInfo>
        <Name Type="LGL">
          <FamilyName>Brown</FamilyName>
          <GivenName>James</GivenName>
        </Name>
        <Address Type="0123" Role="2382">
          <Street>
            <Line1>23 E. 43rd St.</Line1>
          </Street>
          <City>Chicago</City>
         <StateProvince>IL</StateProvince>
          <Country>8104</Country>
          <PostalCode>60611</PostalCode>
        </Address>
        <EmailList>
          <Email Type="01">jdr@ABC.com</Email>
        </EmailList>
        <PhoneNumberList>
          <PhoneNumber Type="0096">
            <Number>(312) 555-1234</Number>
          </PhoneNumber>
        </PhoneNumberList>
      </ContactInfo>
      <CustomerId>0023556</CustomerId>
      <ABN>56402325367</ABN>
      <RegisteredForGST>Y</RegisteredForGST>
    </VendorInfo>
    

`
